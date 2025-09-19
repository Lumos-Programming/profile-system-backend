package handler

import (
	"cloud.google.com/go/firestore"
	"github.com/Lumos-Programming/profile-system-backend/api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	firestoreCollection = "profiles"
	firestoreDocID      = "default" // 単一プロフィールの場合
)

type Handler struct {
	fs *firestore.Client
}

func NewHandler(f *firestore.Client) *Handler {
	return &Handler{
		fs: f,
	}
}

func (h *Handler) GetApiProfileBasicInfo(c *gin.Context) {
	doc, err := h.fs.Collection(firestoreCollection).Doc(firestoreDocID).Get(c)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			c.JSON(200, api.BasicInfo{})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var info api.BasicInfo
	if err := doc.DataTo(&info); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, info)
}

func (h *Handler) PutApiProfileBasicInfo(c *gin.Context) {
	var req api.BasicInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.fs.Collection(firestoreCollection).Doc(firestoreDocID).Set(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, req)
}
