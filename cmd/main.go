package main

import (
	"GoTestApi/internal/auth"
	"GoTestApi/internal/config"
	"GoTestApi/internal/db"
	"GoTestApi/internal/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設定ファイルの読み込み
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalln("Error loading config file:", err)
	}

	// データベース接続の初期化
	database, err := db.InitDB(cfg.DSN)
	if err != nil {
		log.Fatalln("Error connecting to the database:", err)
	}

	// Ginルーターの設定
	router := gin.Default()

	// 認証ルート
	router.POST("/login", func(c *gin.Context) { auth.Login(c, database) })
	router.POST("/register", func(c *gin.Context) { auth.Register(c, database) })

	// プロフィールルート
	router.GET("/profile", auth.Authenticate, func(c *gin.Context) { user.Profile(c, database) })

	// サーバーの開始
	if err := router.Run(":8080"); err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
