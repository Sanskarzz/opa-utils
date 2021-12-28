package resourcesresults

import (
	"github.com/armosec/opa-utils/reporthandling/apis"
	helpersv1 "github.com/armosec/opa-utils/reporthandling/helpers/v1"
)

// GetID get control ID
func (control *ResourceAssociatedControl) GetID() string {
	return control.ControlID
}

// SetID set control ID
func (control *ResourceAssociatedControl) SetID(id string) {
	control.ControlID = id
}

// =============================== Status ====================================

// Status get control status
func (control *ResourceAssociatedControl) GetStatus(f *helpersv1.Filters) apis.IStatus {
	status := apis.StatusPassed // if len(control.ResourceAssociatedRules) == 0 the resource passed
	for i := range control.ResourceAssociatedRules {
		status = apis.Compare(status, control.ResourceAssociatedRules[i].GetStatus(f).Status())
	}
	return helpersv1.NewStatus(status)
}
