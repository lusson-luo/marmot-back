package dao

import "marmot/internal/dao/internal"

type inspectionDao struct {
	*internal.InspectionDao
}

var (
	// User is globally public accessible object for table gf_user operations.
	Inspection = inspectionDao{
		internal.NewInspectionDao(),
	}
)
