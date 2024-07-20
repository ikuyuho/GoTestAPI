package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Profile(c *gin.Context, db *sqlx.DB) {
	// プロフィール取得のロジックをここに記述
	c.JSON(http.StatusOK, gin.H{"message": "profile information"})
}
