package service

import (
	"gjob-admin/pkg/dao"
	"gjob-admin/pkg/model"
)

type sqljob struct{}

var SqljobService sqljob

// 增删改查sqljob
// Create 创建sqljob
func (s *sqljob) Create(data *model.SqlJob) (err error) {
	err = dao.Sqljob.Create(data)
	if err != nil {
		return err
	}
	return nil

}

// Get 获取sqljob
func (s *sqljob) Get(name string) (sqljoblist []model.SqlJob, err error) {
	sqljoblist, err = dao.Sqljob.Get(name)
	if err != nil {
		return nil, err
	}
	return sqljoblist, nil
}

// GetByPage 分页获取sqljob
func (s *sqljob) GetByPage(pagenum, pagesize string) (sqljoblist []model.SqlJob, err error) {
	sqljoblist, err = dao.Sqljob.GetByPage(pagenum, pagesize)
	if err != nil {
		return nil, err
	}
	return sqljoblist, nil
}

func (s *sqljob) GetAll() (sqljoblist []model.SqlJob, err error) {
	sqljoblist, err = dao.Sqljob.GetAll()
	if err != nil {
		return nil, err
	}
	return sqljoblist, nil
}
