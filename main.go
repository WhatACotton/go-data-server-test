package main

import (
	"unify/internal/database"
	"unify/validation"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	validation.CORS(r)
	validation.SessionConfig(r)
	database.TestSQL()

	r.Run(":8080") // 0.0.0.0:8080 でサーバーを立てます。
}
