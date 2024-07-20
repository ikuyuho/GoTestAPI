package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Register(c *gin.Context, db *sqlx.DB) {
	// ユーザー登録処理のロジックをここに記述
	c.JSON(http.StatusOK, gin.H{"message": "register successful"})
}
