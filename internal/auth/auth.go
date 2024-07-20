package auth

import (
	"github.com/gin-gonic/gin"
	"GoTestApi/pkg/models"
)

func authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")   // Authorizationヘッダーからトークンを取得
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		c.Abort()   // リクエストを中止
		return
	}

	claims := &models.Claims{
		
	}
	// トークンをパースし、クレームを検証
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// トークンが無効またはエラーが発生した場合
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()   // リクエストを中止
		return
	}

	// ユーザー情報をコンテキストに設定し、次のハンドラに渡す
	c.Set("user", claims.Username)
	c.Next()   // 次のハンドラに進む
}

