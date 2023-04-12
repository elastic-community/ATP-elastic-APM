package atp

import (
	"atp/db"

	"github.com/gin-gonic/gin"
)

type Player struct {
	ID       int    `json:"key" gorm:"primaryKey;autoIncrement:true;column:id;`
	FullName string `gorm:"full_name"`
	Point    int    `gorm:"point"`
}

func Ranking(ctx *gin.Context) {

	var players []Player
	DB := db.GetDB(ctx.Request.Context())
	DB = DB.Select("id, full_name, point")
	query := ctx.Request.URL.Query().Get("q")
	if query != "" {
		DB = DB.Where("full_name LIKE ?", "%"+query+"%")
	}
	DB.Find(&players)
	ctx.JSON(200, players)
}
