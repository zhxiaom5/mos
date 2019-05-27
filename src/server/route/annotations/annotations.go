package annotations

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/pkg/e"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// AnnotationList 获取列表
func AnnotationList(ctx *gin.Context) {
	type retData struct {
		ID         uint     `json:"id"`
		IsRegion   bool     `json:"is_region"`
		Project    string   `json:"project"`
		Text       string   `json:"text"`
		Tag        []string `json:"tag"`
		RegionID   uint     `json:"region_id"`
		CreateTime string   `json:"create_time"`
		StartTime  string   `json:"start_time"`
		EndTime    string   `json:"end_time"`
	}
	var (
		aList   []Annotation
		retList []retData
		total   int
	)

	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("current_page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)

	search := ctx.Query("search")

	queryDb := glo.Db.Model(&Annotation{})
	if search != `` {
		// 根据username模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		queryDb = queryDb.Where("text LIKE ?", fmt.Sprintf("%%%s%%", search))
	}
	if err := queryDb.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&aList).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": "get query error, " + err.Error(),
		})
		return
	}
	if len(aList) == 0 {
		retList = []retData{}
	}
	for _, r := range aList {
		tagList := strings.Split(r.Tag, ",")
		retList = append(retList, retData{
			ID:         r.ID,
			IsRegion:   r.IsRegion,
			Project:    r.Project,
			Text:       r.Text,
			Tag:        tagList,
			RegionID:   r.RegionID,
			CreateTime: comfunc.FormatTs(r.CreatedAt.Unix()),
			StartTime:  comfunc.FormatTs(r.StartTime),
			EndTime:    comfunc.FormatTs(r.EndTime),
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
}

// AnnotationAdd 添加AnnotationAdd信息
func AnnotationAdd(ctx *gin.Context) {
	type reqPostData struct {
		IsRegion  bool   `json:"isRegion"`
		Project   string `json:"project"`
		Text      string `json:"text"`
		Tags      string `json:"tags"`
		StartTime int64  `json:"startTime"`
		EndTime   int64  `json:"endTime"`
	}

	var (
		req      reqPostData
		a        Annotation
		regionID uint
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
	if err := glo.Db.Model(&Annotation{}).Last(&a).Error; err != nil {
		regionID = 1
		// ctx.JSON(http.StatusOK, gin.H{
		// 	"code":    e.ERROR,
		// 	"message": e.ERROR_MSG,
		// 	"debug":   err.Error(),
		// })
		// return
	} else {
		regionID = a.RegionID + 1
	}
	if req.IsRegion == false {
		glo.Db.Model(&Annotation{}).Create(&Annotation{
			IsRegion:  false,
			Project:   req.Project,
			Text:      req.Text,
			Tag:       req.Tags,
			RegionID:  regionID,
			StartTime: req.StartTime,
		})
	} else {
		glo.Db.Model(&Annotation{}).Create(&Annotation{
			IsRegion:  true,
			Project:   req.Project,
			Text:      req.Text,
			Tag:       req.Tags,
			RegionID:  regionID,
			StartTime: req.StartTime,
		})
		glo.Db.Model(&Annotation{}).Create(&Annotation{
			IsRegion:  true,
			Project:   req.Project,
			Text:      req.Text,
			Tag:       req.Tags,
			RegionID:  regionID,
			StartTime: req.EndTime,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": e.SUCCESS_MSG,
	})
	return
}

// AnnotationDelete 删除annotation
func AnnotationDelete(ctx *gin.Context) {
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
	if err := glo.Db.Model(&Annotation{}).Unscoped().Where("id = ?", req.ID).Delete(&Annotation{}).Error; err != nil {
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
func Tst(ctx *gin.Context) {
	fmt.Printf("%s", comfunc.FormatTs(time.Now().Unix()))
	time.Sleep(time.Second * 30)
	fmt.Printf("%s", comfunc.FormatTs(time.Now().Unix()))
	ctx.JSON(http.StatusOK, gin.H{
		"code": "1",
	})
	return
}
