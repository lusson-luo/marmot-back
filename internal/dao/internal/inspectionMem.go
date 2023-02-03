package internal

import (
	"marmot/internal/model/entity"
)

type InspectionMemDao struct {
	inspects []*entity.Inspection
}

func NewInspectionMemDao() (dao *InspectionMemDao) {
	dao = &InspectionMemDao{}
	return
}

func (dao *InspectionMemDao) Update(inspect entity.Inspection) bool {
	index := dao.findById(inspect.Id)
	if index == -1 {
		return false
	}
	dao.inspects[index] = &inspect
	return true
}

func (dao *InspectionMemDao) Insert(inspect entity.Inspection) {
	dao.inspects = append(dao.inspects, &inspect)
}

func (dao *InspectionMemDao) Delete(id int) bool {
	index := dao.findById(id)
	if index == -1 {
		return false
	}
	dao.inspects = append(dao.inspects[:index], dao.inspects[index+1:]...)
	return true
}

// 根据id查询切片对应下标，如果没有找到，返回-1
func (dao *InspectionMemDao) findById(id int) (index int) {
	index = -1
	for i, model := range dao.inspects {
		if id == model.Id {
			index = i
		}
	}
	return
}

func (dao *InspectionMemDao) FindById(id int) (inspection entity.Inspection, exist bool) {
	index := dao.findById(id)
	if index == -1 {
		exist = false
		return
	}
	inspection = *dao.inspects[index]
	exist = true
	return
}

func (dao *InspectionMemDao) FindAll() (inspects []*entity.Inspection) {
	inspects = dao.inspects
	return
}
