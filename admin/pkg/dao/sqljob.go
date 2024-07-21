package dao

import "gjob-admin/pkg/model"

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
func (s sqljob) GetAll(page, limit int) (sqljoblist []model.SqlJob, err error) {
	rec := DB.Limit(limit).Offset((page - 1) * limit).Find(&sqljoblist)
	return sqljoblist, rec.Error
}
