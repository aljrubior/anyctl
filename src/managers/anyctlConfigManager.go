package managers

import (
	"fmt"
	"github.com/aljrubior/anyctl/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func NewAnyctlConfigManager(appConfig *conf.AppConfig) (*AnyctlConfigManager, error) {
	newManager := AnyctlConfigManager{
		appConfig: appConfig,
	}

	if err := newManager.loadConfiguration(); err != nil {
		return nil, err
	}

	return &newManager, nil

}

type AnyctlConfigManager struct {
	appConfig    *conf.AppConfig
	anyctlConfig *conf.AnyctlConfig
}

func (this *AnyctlConfigManager) loadConfiguration() error {

	anypointHome, err := this.getConfigDir()

	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", anypointHome, this.appConfig.AnyctlConfigFile())

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &this.anyctlConfig)

	if err != nil {
		return err
	}

	return nil
}

func (this *AnyctlConfigManager) NewDeploymentConfigClient() *conf.DeploymentClientConfig {
	return &conf.DeploymentClientConfig{
		Protocol:                     this.anyctlConfig.Anypoint.Protocol,
		Host:                         this.anyctlConfig.Anypoint.Host,
		Port:                         this.anyctlConfig.Anypoint.Port,
		DeploymentPathTemplate:       this.anyctlConfig.Anypoint.Resources.RuntimeManager.Deployments.DeploymentPath,
		DeploymentsPathTemplate:      this.anyctlConfig.Anypoint.Resources.RuntimeManager.Deployments.DeploymentsPath,
		UpdateDeploymentPathTemplate: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Deployments.UpdateDeploymentPath,
	}
}

func (this *AnyctlConfigManager) NewSchedulerClientConfig() *conf.SchedulerClientConfig {
	return &conf.SchedulerClientConfig{
		Protocol:                 this.anyctlConfig.Anypoint.Protocol,
		Host:                     this.anyctlConfig.Anypoint.Host,
		Port:                     this.anyctlConfig.Anypoint.Port,
		SchedulerPathTemplate:    this.anyctlConfig.Anypoint.Resources.RuntimeManager.Schedulers.SchedulerPath,
		SchedulersPathTemplate:   this.anyctlConfig.Anypoint.Resources.RuntimeManager.Schedulers.SchedulersPath,
		RunSchedulerPathTemplate: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Schedulers.RunSchedulerPath,
	}
}

func (this *AnyctlConfigManager) NewAssetClientConfig() *conf.AssetClientConfig {
	return &conf.AssetClientConfig{
		Protocol:            this.anyctlConfig.Anypoint.Protocol,
		Host:                this.anyctlConfig.Anypoint.Host,
		Port:                this.anyctlConfig.Anypoint.Port,
		AssetsPath:          this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.AssetsPath,
		LatestVersionPath:   this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.LatestVersionPath,
		SpecificVersionPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.SpecificVersionPath,
		UploadAssetPath:     this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.UploadAssetPath,
	}
}

func (this *AnyctlConfigManager) NewTargetClientConfig() *conf.TargetClientConfig {
	return &conf.TargetClientConfig{
		Protocol:    this.anyctlConfig.Anypoint.Protocol,
		Host:        this.anyctlConfig.Anypoint.Host,
		Port:        this.anyctlConfig.Anypoint.Port,
		TargetsPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Targets.TargetsPath,
		FabricsPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Targets.FabricsPath,
	}
}

func (this *AnyctlConfigManager) NewRuntimeFabricClientConfig() *conf.RuntimeFabricClientConfig {
	return &conf.RuntimeFabricClientConfig{
		Protocol:               this.anyctlConfig.Anypoint.Protocol,
		Host:                   this.anyctlConfig.Anypoint.Host,
		Port:                   this.anyctlConfig.Anypoint.Port,
		FabricPath:             this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.FabricPath,
		FabricsPath:            this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.FabricsPath,
		TargetPath:             this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.TargetPath,
		TargetsPath:            this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.TargetsPath,
		PrivateSpacePath:       this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.PrivateSpacePath,
		PrivateSpacesPath:      this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.PrivateSpacesPath,
		PrivateSpaceFabricPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.RuntimeFabrics.PrivateSpaceFabricPath,
	}
}

func (this *AnyctlConfigManager) NewSharedSpaceClientConfig() *conf.SharedSpaceClientConfig {
	return &conf.SharedSpaceClientConfig{
		Protocol:         this.anyctlConfig.Anypoint.Protocol,
		Host:             this.anyctlConfig.Anypoint.Host,
		Port:             this.anyctlConfig.Anypoint.Port,
		SharedSpacePath:  this.anyctlConfig.Anypoint.Resources.Admin.SharedSpaces.SharedSpacePath,
		SharedSpacesPath: this.anyctlConfig.Anypoint.Resources.Admin.SharedSpaces.SharedSpacesPath,
	}
}

func (this *AnyctlConfigManager) NewPrivateSpaceClientConfig() *conf.PrivateSpaceClientConfig {
	return &conf.PrivateSpaceClientConfig{
		Protocol:          this.anyctlConfig.Anypoint.Protocol,
		Host:              this.anyctlConfig.Anypoint.Host,
		Port:              this.anyctlConfig.Anypoint.Port,
		PrivateSpacePath:  this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.PrivateSpacePath,
		PrivateSpacesPath: this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.PrivateSpacesPath,
		FabricPath:        this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.FabricPath,
		FabricsPath:       this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.FabricsPath,
	}
}

func (this *AnyctlConfigManager) NewFabricClientConfig() *conf.FabricClientConfig {
	return &conf.FabricClientConfig{
		Protocol:    this.anyctlConfig.Anypoint.Protocol,
		Host:        this.anyctlConfig.Anypoint.Host,
		Port:        this.anyctlConfig.Anypoint.Port,
		FabricPath:  this.anyctlConfig.Anypoint.Resources.Admin.Fabrics.FabricPath,
		FabricsPath: this.anyctlConfig.Anypoint.Resources.Admin.Fabrics.FabricsPath,
	}
}

func (this *AnyctlConfigManager) NewAccountClientConfig() *conf.AccountClientConfig {
	return &conf.AccountClientConfig{
		Protocol:                 this.anyctlConfig.Anypoint.Protocol,
		Host:                     this.anyctlConfig.Anypoint.Host,
		Port:                     this.anyctlConfig.Anypoint.Port,
		LoginPath:                this.anyctlConfig.Anypoint.Resources.Admin.Accounts.LoginPath,
		OrganizationPathTemplate: this.anyctlConfig.Anypoint.Resources.Admin.Accounts.OrganizationPath,
	}
}

func (this AnyctlConfigManager) getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", homeDir, this.appConfig.ConfigDir()), nil
}
