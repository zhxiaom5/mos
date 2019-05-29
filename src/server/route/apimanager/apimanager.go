package apimanager

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/pkg/e"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// APIList 获取列表
func APIList(ctx *gin.Context) {
	var (
		retList []APITabInfo
		aList   []APITab
		total   int
	)

	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("current_page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)

	search := ctx.Query("search")

	queryDb := glo.Db.Set("gorm:auto_preload", true).Model(&APITab{})
	if search != `` {
		// 根据api_uri模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		queryDb = queryDb.Where("api_uri LIKE ?", fmt.Sprintf("%%%s%%", search))
	}
	if err := queryDb.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&aList).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "get query error, " + err.Error(),
		})
		return
	}
	if len(aList) == 0 {
		retList = []APITabInfo{}
	}
	for _, r := range aList {
		tmpAPIKeyInfo := []APIKeyInfo{}
		if len(r.APIKey) > 0 {
			for _, m := range r.APIKey {
				tmpAPIKeyInfo = append(tmpAPIKeyInfo, APIKeyInfo{
					ID:        m.ID,
					AuthKey:   m.AuthKey,
					User:      m.User,
					IsDisable: m.IsDisable,
				})
			}
		}
		retList = append(retList, APITabInfo{
			ID:         r.ID,
			APIURI:     r.APIURI,
			Content:    r.Content,
			APIKeyInfo: tmpAPIKeyInfo,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":       e.SUCCESS,
		"message":    e.SUCCESS_MSG,
		"data":       retList,
		"total":      total,
		"total_page": comfunc.FlorPageInt(pageSize, total),
		"page_size":  pageSize,
		"page":       page,
	})
	return
}

// APIAdd 添加Api信息
func APIAdd(ctx *gin.Context) {
	type reqPostData struct {
		APIURI  string `json:"uri"`
		Content string `json:"content"`
	}

	var (
		req reqPostData
	)

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
			"debug":   err.Error(),
			"data":    req,
		})
		return
	}
	a := APITab{
		APIURI:  req.APIURI,
		Content: req.Content,
	}
	if err := glo.Db.Model(&APITab{}).Create(&a).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG,
			"debug":   fmt.Sprintf("添加失败: %s", err.Error()),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}

// APIDelete 删除API
func APIDelete(ctx *gin.Context) {
	type reqData struct {
		ID uint `json:"id"`
	}
	var (
		req reqData
	)
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	if err := glo.Db.Model(&APITab{}).Unscoped().Where("id = ?", req.ID).Delete(&APITab{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "删除失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "删除成功",
	})
	return
}

// APIKeyAdd 添加API信息
func APIKeyAdd(ctx *gin.Context) {
	type reqPostData struct {
		APIID        uint   `json:"apiID"`
		User         string `json:"user"`
		IsKeyDisable bool   `json:"isKeyDisable"`
	}

	var (
		req reqPostData
	)

	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
			"debug":   err.Error(),
			"data":    req,
		})
		return
	}

	a := APIKey{
		APITabID:  req.APIID,
		User:      req.User,
		IsDisable: req.IsKeyDisable,
		AuthKey:   comfunc.Md5String(comfunc.FormatTs(time.Now().Unix())),
	}
	if err := glo.Db.Model(&APIKey{}).Create(&a).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_MSG,
			"debug":   fmt.Sprintf("添加失败: %s", err.Error()),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}

// APIKeyDelete 删除API KEY
func APIKeyDelete(ctx *gin.Context) {
	type reqData struct {
		ID uint `json:"id"`
	}
	var (
		req reqData
	)
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": e.PARAM_ERROR_MSG,
		})
		return
	}
	if err := glo.Db.Model(&APIKey{}).Unscoped().Where("id = ?", req.ID).Delete(&APIKey{}).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "删除失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "删除成功",
	})
	return
}
