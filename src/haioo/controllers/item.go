package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"haioo/code"
	"haioo/model"
	"haioo/utils"
	"time"
)

type ItemController struct {
	BaseController
}

//获取广告位数据
func (this ItemController) GetUseableAd(c *gin.Context) {
	//获取外部参数
	c.Request.ParseForm()
	adPostId := utils.ParseInt(c.Request.Form.Get("adPostId"))
	num := utils.ParseInt(c.Request.Form.Get("num"))
	//验证参数合法性
	if adPostId <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "adPostId必须大于0", map[string]interface{}{})
		return
	}
	if num <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "num必须大于0", map[string]interface{}{})
		return
	}
	//查询数据库
	advsCon := make([]model.HaiouAdvsCon, 0)
	nowTime := time.Now().Unix()
	model.XormRead.Where("group_id=? AND stime<=? AND etime>=? AND isopen=1", adPostId, nowTime, nowTime).OrderBy("`create_time` DESC").Limit(4, 0).Find(&advsCon)
	this.ResponseData(c, code.Response.SUCCESS, "", advsCon)
}

//获取活动列表数据
func (this ItemController) GetActivityList(c *gin.Context) {
	//获取外部参数
	c.Request.ParseForm()
	catTagId := utils.ParseInt(c.Request.Form.Get("catTagId"))
	userId := utils.ParseInt(c.Request.Form.Get("userId"))
	currentPage := utils.ParseInt(c.Request.Form.Get("currentPage"))
	pageSize := utils.ParseInt(c.Request.Form.Get("pageSize"))

	//检查参数
	if catTagId <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "catTagId参数必须大于0", nil)
		return
	}
	if userId < 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "uesrId参数必须大于等于0", nil)
		return
	}
	if currentPage <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "currentPage参数必须大于0", nil)
		return
	}
	if pageSize <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "pageSize参数必须大于0", nil)
		return
	}

	//组装条件，获取数据
	baseWhere := ""
	if catTagId > 0 {
		baseWhere = fmt.Sprintf(" AND LOCATE('%d',CONCAT(',',`tags`))>0", catTagId)
	}
	nowTime := time.Now().Unix()
	activityList := make([]model.HaiouActivity, 0)
	model.XormRead.Cols("id", "title", "pic", "discount", "start_time", "end_time", "catid", "brandid").
		Where(fmt.Sprintf("status>0 AND end_time>%d %s", nowTime, baseWhere)).
		OrderBy(fmt.Sprintf("IF(`start_time`<=%d AND `end_time`>=%d,1,0) DESC,`displayorder` DESC,IF(`start_time`<=%d AND `end_time`>=%d,`start_time`,0-`start_time`) DESC", nowTime, nowTime, nowTime, nowTime)).
		Limit(pageSize, currentPage-1).
		Find(&activityList)
	activityListMap := make([]map[string]interface{}, 0)
	listLength := len(activityList)
	if listLength > 0 {
		//标记活动数据是否开始,加入当前时间
		activityIds := ""
		for _, v := range activityList {
			itemMap := utils.StructToMap(v)
			itemMap["nowTime"] = nowTime
			//是否开始
			endTime, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%d", v.EndTime))
			startTime, _ := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%d", v.StartTime))
			if nowTime > startTime.Unix() && nowTime < endTime.Unix() {
				itemMap["IsStarting"] = 1
			} else {
				itemMap["IsStarting"] = 0
			}
			activityListMap = append(activityListMap, itemMap)
			if activityIds != "" {
				activityIds += ","
			}
			activityIds += fmt.Sprintf("%d", v.Id)
		}
		//标记活动是否收藏
		favoriteBrandsMap := make(map[int]bool)
		if userId > 0 {
			favoriteBrands := make([]model.HaiouMemberFavoriteBrand, 0)
			model.XormRead.Cols("activity_id").Where("member_id=? AND activity_id IN (?) AND status>0", userId, activityIds).Limit(listLength).Find(&favoriteBrands)
			if len(favoriteBrands) > 0 {
				for _, v := range favoriteBrands {
					favoriteBrandsMap[v.ActivityId] = true
				}
			}
		}
		for _, v := range activityListMap {
			activityId := v["Id"].(int)
			if _, ok := favoriteBrandsMap[activityId]; ok {
				v["IsFavorite"] = 1
			} else {
				v["IsFavorite"] = 0
			}
		}
	}
	this.ResponseData(c, code.Response.SUCCESS, "", activityListMap)
}

//获取分类列表
func (this ItemController) GetCatTags(c *gin.Context) {
	//获取外部参数
	c.Request.ParseForm()
	parentId := utils.ParseInt(c.Request.Form.Get("parentId"))

	//检查外部参数
	if parentId < 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "无效的parentId", nil)
	}
	where := ""
	if parentId == 0 {
		where = "`status`>0 AND`id`>=1000 AND `id`<=9999 "
	} else {
		where = fmt.Sprintf("`status`>0 AND `id` LIKE '%d__')", parentId)
	}

	//查询数据
	catTags := make([]model.HaiouProductTags, 0)
	model.XormRead.Cols("id", "tags", "logo").Where(where).OrderBy("`displayorder` ASC").Find(&catTags)
	this.ResponseData(c, code.Response.SUCCESS, "", catTags)
}

//添加特卖收藏
func (this ItemController) AddActivityToCollect(c *gin.Context) {
	//获取外部参数
	c.Request.ParseForm()
	userId := utils.ParseInt(c.Request.Form.Get("userId"))
	activityId := utils.ParseInt(c.Request.Form.Get("activityId"))
	clientPlatform := utils.ParseInt(c.Request.Form.Get("clientPlatform"))
	clientSystem := utils.ParseInt(c.Request.Form.Get("clientSystem"))

	//检查外部参数
	if userId <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "userId必须大于0", nil)
	}
	if activityId <= 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "activityId必须大于0", nil)
	}
	if clientPlatform < 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "clientPlatform不合法", nil)
	}
	if clientSystem < 0 {
		this.ResponseData(c, code.Response.INVALID_PARAMETER, "clientSystem不合法", nil)
	}

}
