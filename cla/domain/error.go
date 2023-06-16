package domain

import "github.com/opensourceways/app-cla-signing/cla/domain/dp"

var (
	errorNotSameCorp              = dp.NewDomainError("not_same_corp")
	errorEmailDomainExists        = dp.NewDomainError("email_domain_exists")
	errorErrAdminAsManager        = dp.NewDomainError("admin_as_manager")
	errorUnmatchedEmailDomain     = dp.NewDomainError("unmatched_email_domain")
	errorRemoveEnabledSigning     = dp.NewDomainError("remove_enabled_signing")
	errorEmployeeManagerExists    = dp.NewDomainError("corp_manager_exists")
	errorTooManyEmployeeManagers  = dp.NewDomainError("many_employee_managers")
	errorEmployeeManagerNotExists = dp.NewDomainError("corp_manager_not_exists")
)
