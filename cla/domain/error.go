package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

var (
	errorNotSameCorp              = dp.NewDomainError("not_same_corp")
	errorErrAdminAsManager        = dp.NewDomainError("admin_as_manager")
	errorEmployeeManagerExists    = dp.NewDomainError("corp_manager_exists")
	errorTooManyEmployeeManagers  = dp.NewDomainError("many_employee_managers")
	errorEmployeeManagerNotExists = dp.NewDomainError("corp_manager_not_exists")
)
