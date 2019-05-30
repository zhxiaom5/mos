package jenkins

import (
	"fmt"
	"mos/src/glo"
	"mos/src/glo/comfunc"
	"mos/src/pkg/e"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// JenkinsPost jenkins 构建数据上报
func JenkinsPost(ctx *gin.Context) {
	var (
		req JobPost
		j   JenkinsJob
	)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "参数错误," + fmt.Sprintf("错误信息: %s", err),
		})
		return
	}
	// fmt.Printf("%+vn", req)
	if status, err := j.Create(&req); status == false {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.PARAM_ERROR,
			"message": "上报失败," + fmt.Sprintf("错误信息: %s", err),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "上报成功",
	})
	return
}

// JenkinsJobList 获取列表
func JenkinsJobList(ctx *gin.Context) {
	type retData struct {
		ID          uint   `json:"id"`
		Project     string `json:"project"`
		Module      string `json:"module"`
		Tag         string `json:"tag"`
		Title       string `json:"title"`
		Hosts       string `json:"hosts"`
		BuildStatus string `json:"build_status"`
		BuildUser   string `json:"build_user"`
		CreateTime  string `json:"create_time"`
	}
	var (
		aList   []JenkinsJob
		retList []retData
		total   int
	)

	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	page, _ := strconv.Atoi(ctx.Query("current_page"))
	pageSize = comfunc.GetDefaultPageSize(pageSize)
	page = comfunc.GetDefaultPage(page)

	search := ctx.Query("search")

	queryDb := glo.Db.Model(&JenkinsJob{})
	if search != `` {
		// 根据username模糊查询，可以将gorm链接添加条件后，赋值覆盖自身，得到不定条件的链式查询效果
		queryDb = queryDb.Where("project LIKE ?", fmt.Sprintf("%%%s%%", search))
	}
	if err := queryDb.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&aList).Error; err != nil {
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
		retList = append(retList, retData{
			ID:          r.ID,
			Project:     r.Project,
			Module:      r.Module,
			Title:       r.Title,
			Tag:         r.Tag,
			Hosts:       r.Hosts,
			CreateTime:  comfunc.FormatTs(r.CreatedAt.Unix()),
			BuildStatus: r.BuildStatus,
			BuildUser:   r.BuildUser,
		})
	}
	queryDb.Count(&total)
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
