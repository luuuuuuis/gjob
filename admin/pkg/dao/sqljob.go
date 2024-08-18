package dao

import (
	"gjob-admin/pkg/model"
	"strconv"
)

var Sqljob sqljob

type sqljob struct{}

// 定义sqljob结构体的增删改查功能
// 增
func (s sqljob) Create(sqljob *model.SqlJob) (err error) {
	err = DB.Create(sqljob).Error
	return
}

// 删
// 根据任务名删除
func (s sqljob) Delete(name string) (err error) {
	// 找到名字为name的记录然后删除
	err = DB.Where("name = ?", name).First(&model.SqlJob{}).Delete(&model.SqlJob{}).Error
	return
}

// 改
func (s sqljob) Update(sqljob *model.SqlJob) (err error) {
	err = DB.Save(sqljob).Error
	return
}

// 查
// 根据任务名模糊查
func (s sqljob) Get(name string) (sqljoblist []model.SqlJob, err error) {
	rec := DB.Where("name LIKE ?", name).Find(&sqljoblist)
	return sqljoblist, rec.Error
}

// 分页查询全部
func (s sqljob) GetByPage(pagenum, pagesize string) (sqljoblist []model.SqlJob, err error) {
	pageSize, offsetVal := PageComput(pagenum, pagesize)
	rec := DB.Limit(pageSize).Offset(offsetVal).Find(&sqljoblist)
	return sqljoblist, rec.Error
}

// 分页计算
func PageComput(pagenum, pagesize string) (pageSize, offsetVal int) {
	// 判断是否有数据
	pageSize, _ = strconv.Atoi(pagesize)
	pageNum, _ := strconv.Atoi(pagenum)
	// 判断是否需要分页
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	offsetVal = (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}
	return
}

// 全量查询
func (s sqljob) GetAll() (sqljoblist []model.SqlJob, err error) {
	rec := DB.Find(&sqljoblist)
	return sqljoblist, rec.Error
}

