package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BasicInfo struct {
	StudentID        string `json:"student_id"`
	Faculty          string `json:"faculty"`
	LastName         string `json:"last_name"`
	FirstName        string `json:"first_name"`
	Nickname         string `json:"nickname"`
	SelfIntroduction string `json:"self_introduction"`
}

const firestoreCollection = "profiles"
const firestoreDocID = "default" // 単一プロフィールの場合

func main() {
	cfg, err := loadConfig()
	if err != nil {
		slog.Error("Config load error: %v", err)
		os.Exit(1)
	}

	opts := option.WithCredentialsFile(cfg.Firestore.Credentials)
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, cfg.Firestore.ProjectID, opts)
	if err != nil {
		slog.Error("Firestore client error: %v", err)
		os.Exit(1)
	}

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// プロフィール基本情報取得
	router.GET("/api/profile/basic-info", func(c *gin.Context) {
		doc, err := client.Collection(firestoreCollection).Doc(firestoreDocID).Get(ctx)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				c.JSON(200, BasicInfo{})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var info BasicInfo
		if err := doc.DataTo(&info); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, info)
	})

	// プロフィール基本情報更新
	router.PUT("/api/profile/basic-info", func(c *gin.Context) {
		var req BasicInfo
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		_, err := client.Collection(firestoreCollection).Doc(firestoreDocID).Set(ctx, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, req)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen: %s\n", err)
			os.Exit(1)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server...")
	client.Close()
}
