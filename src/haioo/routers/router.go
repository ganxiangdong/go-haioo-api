package routers

import (
	"github.com/gin-gonic/gin"
	"haioo/controllers"
)

func Init(r *gin.Engine) {
	//版本号1
	v1 := r.Group("/v1")
	v1.GET("/item/getUseableAd", controllers.ItemController{}.GetUseableAd)
	v1.GET("/item/getActivityList", controllers.ItemController{}.GetActivityList)
	v1.GET("/item/getCatTags", controllers.ItemController{}.GetCatTags)
	v1.POST("/item/AddActivityToCollect", controllers.ItemController{}.AddActivityToCollect)

}
