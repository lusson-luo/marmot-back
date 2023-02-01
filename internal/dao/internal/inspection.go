package internal

import (
	"marmot/internal/model/do"
)

type InspectionDao struct {
	inspects []*do.Inspection
}

func NewInspectionDao() (dao *InspectionDao) {
	dao = &InspectionDao{}
	return
}

func (dao *InspectionDao) Update(inspect do.Inspection) bool {
	index := dao.findById(inspect.Id)
	if index == -1 {
		return false
	}
	dao.inspects[index] = &inspect
	return true
}

func (dao *InspectionDao) Insert(inspect do.Inspection) {
	dao.inspects = append(dao.inspects, &inspect)
}

func (dao *InspectionDao) Delete(id int) bool {
	index := dao.findById(id)
	if index == -1 {
		return false
	}
	dao.inspects = append(dao.inspects[:index], dao.inspects[index+1:]...)
	return true
}

// 根据id查询切片对应下标，如果没有找到，返回-1
func (dao *InspectionDao) findById(id int) (index int) {
	index = -1
	for i, model := range dao.inspects {
		if id == model.Id {
			index = i
		}
	}
	return
}

func (dao *InspectionDao) FindById(id int) (inspection do.Inspection, exist bool) {
	index := dao.findById(id)
	if index == -1 {
		exist = false
		return
	}
	inspection = *dao.inspects[index]
	exist = true
	return
}

func (dao *InspectionDao) FindAll() (inspects []*do.Inspection) {
	inspects = dao.inspects
	return
}
