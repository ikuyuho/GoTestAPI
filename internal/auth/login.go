package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Login(c *gin.Context, db *sqlx.DB) {
	// ログイン処理のロジックをここに記述
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
