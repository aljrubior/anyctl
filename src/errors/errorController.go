package errors

import (
	"fmt"
	"github.com/aljrubior/anyctl/errors/advisors"
)

func Catch(err error) *errorController {
	return &errorController{
		err: err,
	}
}

type errorController struct {
	err error
}

func (this errorController) Println() {

	switch this.err.(type) {
	case *UnauthorizedError:
		advisors.NewUnauthorizedAdviser(this.err).Advise()
	case *DeploymentNotFoundError:
		if error, ok := this.err.(*DeploymentNotFoundError); ok {
			advisors.NewDeploymentNotFoundAdvisor(error.Error(), error.DeploymentName, error.Options).Advise()
		}
	case *TargetNotFoundError:
		if error, ok := this.err.(*TargetNotFoundError); ok {
			advisors.NewTargetNotFoundAdvisor(error.Error(), error.TargetName, error.Options).Advise()
		}
	case *RtfTargetNotFoundError:
		if error, ok := this.err.(*RtfTargetNotFoundError); ok {
			advisors.NewRtfTargetNotFoundAdvisor(error.Error(), error.TargetName, error.Options).Advise()
		}
	case *AssetNotFoundError:
		if error, ok := this.err.(*AssetNotFoundError); ok {
			advisors.NewAssetNotFoundAdvisor(error.Error(), error.AssetName, error.Options).Advise()
		}
	case *SchedulerNotFoundError:
		if error, ok := this.err.(*SchedulerNotFoundError); ok {
			advisors.NewSchedulerNotFoundAdvisor(error.Error(), error.FlowName, error.Options).Advise()
		}
	case *AnypointEnvironmentNotFoundError:
		if error, ok := this.err.(*AnypointEnvironmentNotFoundError); ok {
			advisors.NewAnypointEnvironmentNotFoundAdvisor(error.Error(), error.EnvironmentName, error.Options).Advise()
		}
	case *OrganizationFabricNotFoundError:
		if error, ok := this.err.(*OrganizationFabricNotFoundError); ok {
			advisors.NewOrganizationFabricNotFoundAdvisor(error.Error(), error.RuntimeFabricName, error.Options).Advise()
		}
	case *AssetGroupNotFoundError:
		if error, ok := this.err.(*AssetGroupNotFoundError); ok {
			advisors.NewAssetGroupNotFoundAdvisor(error.Error(), error.GroupId, error.AssetName, error.Options).Advise()
		}
	default:
		println(fmt.Sprintf("ERROR: %s", this.err))
	}
}
