package managers

import (
	"fmt"
	"github.com/aljrubior/anyctl/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func NewAnyctlConfigManager(appConfig conf.AppConfig) (AnyctlConfigManager, error) {
	newManager := AnyctlConfigManager{
		appConfig: appConfig,
	}

	if err := newManager.loadConfiguration(); err != nil {
		return AnyctlConfigManager{}, err
	}

	return newManager, nil

}

type AnyctlConfigManager struct {
	appConfig    conf.AppConfig
	anyctlConfig conf.AnyctlConfig

	assetClientConfig               *conf.AssetClientConfig
	deploymentClientConfig          *conf.DeploymentClientConfig
	TargetClientConfig              *conf.TargetClientConfig
	OrganizationRuntimeFabricConfig *conf.RuntimeFabricClientConfig
	SharedSpaceClientConfig         *conf.SharedSpaceClientConfig
	privateSpaceClientConfig        *conf.PrivateSpaceClientConfig
	fabricClientConfig              *conf.FabricClientConfig
	schedulerClientConfig           *conf.SchedulerClientConfig
	accountClientConfig             *conf.AccountClientConfig
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

func (this *AnyctlConfigManager) GetDeploymentConfigClient() conf.DeploymentClientConfig {
	if this.deploymentClientConfig == nil {
		this.deploymentClientConfig = &conf.DeploymentClientConfig{
			Protocol:                     this.anyctlConfig.Anypoint.Protocol,
			Host:                         this.anyctlConfig.Anypoint.Host,
			Port:                         this.anyctlConfig.Anypoint.Port,
			DeploymentPathTemplate:       this.anyctlConfig.Anypoint.Resources.RuntimeManager.Deployments.DeploymentPath,
			DeploymentsPathTemplate:      this.anyctlConfig.Anypoint.Resources.RuntimeManager.Deployments.DeploymentsPath,
			UpdateDeploymentPathTemplate: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Deployments.UpdateDeploymentPath,
		}
	}

	return *this.deploymentClientConfig
}

func (this *AnyctlConfigManager) GetSchedulerClientConfig() conf.SchedulerClientConfig {
	if this.schedulerClientConfig == nil {
		this.schedulerClientConfig = &conf.SchedulerClientConfig{
			Protocol:         this.anyctlConfig.Anypoint.Protocol,
			Host:             this.anyctlConfig.Anypoint.Host,
			Port:             this.anyctlConfig.Anypoint.Port,
			SchedulerPath:    this.anyctlConfig.Anypoint.Resources.RuntimeManager.Schedulers.SchedulerPath,
			SchedulersPath:   this.anyctlConfig.Anypoint.Resources.RuntimeManager.Schedulers.SchedulersPath,
			RunSchedulerPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Schedulers.RunSchedulerPath,
		}
	}

	return *this.schedulerClientConfig
}

func (this *AnyctlConfigManager) GetAssetClientConfig() conf.AssetClientConfig {

	if this.assetClientConfig == nil {
		this.assetClientConfig = &conf.AssetClientConfig{
			Protocol:            this.anyctlConfig.Anypoint.Protocol,
			Host:                this.anyctlConfig.Anypoint.Host,
			Port:                this.anyctlConfig.Anypoint.Port,
			AssetsPath:          this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.AssetsPath,
			LatestVersionPath:   this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.LatestVersionPath,
			SpecificVersionPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.SpecificVersionPath,
			UploadAssetPath:     this.anyctlConfig.Anypoint.Resources.RuntimeManager.Assets.UploadAssetPath,
		}
	}

	return *this.assetClientConfig

}

func (this *AnyctlConfigManager) GetTargetClientConfig() conf.TargetClientConfig {
	if this.TargetClientConfig == nil {
		this.TargetClientConfig = &conf.TargetClientConfig{
			Protocol:    this.anyctlConfig.Anypoint.Protocol,
			Host:        this.anyctlConfig.Anypoint.Host,
			Port:        this.anyctlConfig.Anypoint.Port,
			TargetsPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Targets.TargetsPath,
			FabricsPath: this.anyctlConfig.Anypoint.Resources.RuntimeManager.Targets.FabricsPath,
		}
	}

	return *this.TargetClientConfig

}

func (this *AnyctlConfigManager) GetOrganizationRuntimeFabricClientConfig() conf.RuntimeFabricClientConfig {

	if this.OrganizationRuntimeFabricConfig == nil {
		this.OrganizationRuntimeFabricConfig = &conf.RuntimeFabricClientConfig{
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

	return *this.OrganizationRuntimeFabricConfig

}

func (this *AnyctlConfigManager) GetSharedSpaceClientConfig() conf.SharedSpaceClientConfig {
	if this.SharedSpaceClientConfig == nil {
		this.SharedSpaceClientConfig = &conf.SharedSpaceClientConfig{
			Protocol:         this.anyctlConfig.Anypoint.Protocol,
			Host:             this.anyctlConfig.Anypoint.Host,
			Port:             this.anyctlConfig.Anypoint.Port,
			SharedSpacePath:  this.anyctlConfig.Anypoint.Resources.Admin.SharedSpaces.SharedSpacePath,
			SharedSpacesPath: this.anyctlConfig.Anypoint.Resources.Admin.SharedSpaces.SharedSpacesPath,
		}
	}

	return *this.SharedSpaceClientConfig
}

func (this *AnyctlConfigManager) GetPrivateSpaceClientConfig() conf.PrivateSpaceClientConfig {
	if this.privateSpaceClientConfig == nil {
		this.privateSpaceClientConfig = &conf.PrivateSpaceClientConfig{
			Protocol:          this.anyctlConfig.Anypoint.Protocol,
			Host:              this.anyctlConfig.Anypoint.Host,
			Port:              this.anyctlConfig.Anypoint.Port,
			PrivateSpacePath:  this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.PrivateSpacePath,
			PrivateSpacesPath: this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.PrivateSpacesPath,
			FabricPath:        this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.FabricPath,
			FabricsPath:       this.anyctlConfig.Anypoint.Resources.Admin.PrivateSpaces.FabricsPath,
		}
	}

	return *this.privateSpaceClientConfig
}

func (this AnyctlConfigManager) GetFabricClientConfig() conf.FabricClientConfig {

	if this.fabricClientConfig == nil {
		this.fabricClientConfig = &conf.FabricClientConfig{
			Protocol:    this.anyctlConfig.Anypoint.Protocol,
			Host:        this.anyctlConfig.Anypoint.Host,
			Port:        this.anyctlConfig.Anypoint.Port,
			FabricPath:  this.anyctlConfig.Anypoint.Resources.Admin.Fabrics.FabricPath,
			FabricsPath: this.anyctlConfig.Anypoint.Resources.Admin.Fabrics.FabricsPath,
		}
	}

	return *this.fabricClientConfig
}

func (this *AnyctlConfigManager) GetAccountClientConfig() conf.AccountClientConfig {
	if this.accountClientConfig == nil {
		this.accountClientConfig = &conf.AccountClientConfig{
			Protocol:                 this.anyctlConfig.Anypoint.Protocol,
			Host:                     this.anyctlConfig.Anypoint.Host,
			Port:                     this.anyctlConfig.Anypoint.Port,
			LoginPath:                this.anyctlConfig.Anypoint.Resources.Admin.Accounts.LoginPath,
			OrganizationPathTemplate: this.anyctlConfig.Anypoint.Resources.Admin.Accounts.OrganizationPath,
		}
	}

	return *this.accountClientConfig
}

func (this AnyctlConfigManager) getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", homeDir, this.appConfig.ConfigDir()), nil
}
