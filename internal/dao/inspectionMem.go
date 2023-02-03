package dao

import "marmot/internal/dao/internal"

type inspectionDao struct {
	*internal.InspectionMemDao
	*internal.InspectionDao
}

var (
	// User is globally public accessible object for table gf_user operations.
	Inspection = inspectionDao{
		internal.NewInspectionMemDao(),
		internal.NewInspectionDao(),
	}
)
