package cmd

import (
	"fmt"
	"github.com/aljrubior/anyctl/clients/accounts"
	"github.com/aljrubior/anyctl/clients/assets"
	"github.com/aljrubior/anyctl/clients/deployments"
	"github.com/aljrubior/anyctl/clients/fabrics"
	"github.com/aljrubior/anyctl/clients/organizationPrivateSpaces"
	"github.com/aljrubior/anyctl/clients/organizationRuntimeFabrics"
	"github.com/aljrubior/anyctl/clients/privateSpaces"
	"github.com/aljrubior/anyctl/clients/schedulers"
	"github.com/aljrubior/anyctl/clients/sharedspaces"
	"github.com/aljrubior/anyctl/clients/targets"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/services"
	"github.com/aljrubior/anyctl/utils"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var Console utils.Console

var AnyctlConfig *conf.AnyctlConfig

var ConfigManager *managers.DefaultConfigManager
var AssetManager *managers.DefaultAssetManager
var DeploymentManager *managers.DefaultDeploymentManager
var TargetManager *managers.DefaultTargetManager
var OrganizationRutimeFabricManager *managers.DefaultOrganizationRuntimeFabricManager
var OrganizationPrivateSpaceManager *managers.DefaultOrganizationPrivateSpaceManager
var SharedSpaceManager *managers.DefaultSharedSpaceManager
var PrivateSpaceManager *managers.DefaultPrivateSpaceManager
var FabricManager *managers.DefaultFabricManager
var DeployerManager *managers.DefaultDeployerManager
var SchedulerManager *managers.DefaultSchedulerManager
var AccountManager *managers.DefaultAccountManager

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "src",
	Short: "src manages the Runtime Fabric deployments on Anypoint platform",
	//	Long: `A longer description that spans multiple lines and likely contains
	//examples and usage of using your application. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	// Anyctl Config
	appConfig := conf.NewAppConfig()
	anyctlConfigManager, err := managers.NewAnyctlConfigManager(appConfig)

	if err != nil {
		exit(err)
	}

	// Config Injection
	ConfigManager = managers.NewDefaultConfigManager(*appConfig)

	// Asset Injection
	assetConfig := anyctlConfigManager.NewAssetClientConfig()
	assetClient := assets.NewDefaultAssetClient(*assetConfig)
	assetService := services.NewDefaultAssetService(assetClient)
	AssetManager = managers.NewDefaultAssetManager(assetService)

	// Deployment Injection
	deploymentClientConfig := anyctlConfigManager.NewDeploymentConfigClient()
	deploymentClient := deployments.NewDefaultDeploymentClient(*deploymentClientConfig)
	deploymentService := services.NewDefaultDeploymentService(deploymentClient)
	DeploymentManager = managers.NewDefaultDeploymentManager(deploymentService, AssetManager)

	// Target Injection
	targetClientConfig := anyctlConfigManager.NewTargetClientConfig()
	targetClient := targets.NewDefaultTargetClient(*targetClientConfig)
	targetService := services.NewDefaultTargetService(targetClient)
	TargetManager = managers.NewDefaultTargetManager(targetService)

	// RtfTarget Injection
	runtimeFabricClientConfig := anyctlConfigManager.NewRuntimeFabricClientConfig()
	organizationRuntimeFabricClient := organizationRuntimeFabrics.NewDefaultOrganizationRuntimeFabricClient(*runtimeFabricClientConfig)
	organizationRuntimeFabricService := services.NewDefaultOrganizationRuntimeFabricService(organizationRuntimeFabricClient)
	OrganizationRutimeFabricManager = managers.NewDefaultOrganizationRuntimeFabricManager(organizationRuntimeFabricService)

	// Organization Private Space Injection
	organizationPrivateSpaceClient := organizationPrivateSpaces.NewOrganizationDefaultPrivateSpaceClient(*runtimeFabricClientConfig)
	organizationPrivateSpaceService := services.NewDefaultOrganizationPrivateSpaceService(organizationPrivateSpaceClient)
	OrganizationPrivateSpaceManager = managers.NewDefaultOrganizationPrivateSpaceManager(organizationPrivateSpaceService)

	// Shared Space Injection
	sharedSpaceClientConfig := anyctlConfigManager.NewSharedSpaceClientConfig()
	sharedSpaceClient := sharedspaces.NewDefaultSharedSpaceClient(*sharedSpaceClientConfig)
	sharedSpaceService := services.NewDefaultSharedSpaceService(sharedSpaceClient)
	SharedSpaceManager = managers.NewDefaultSharedSpaceManager(sharedSpaceService)

	// Private Space Injection
	privateSpaceClientConfig := anyctlConfigManager.NewPrivateSpaceClientConfig()
	privateSpaceClient := privateSpaces.NewDefaultPrivateSpaceClient(*privateSpaceClientConfig)
	privateSpaceService := services.NewDefaultPrivateSpaceService(privateSpaceClient)
	PrivateSpaceManager = managers.NewDefaultPrivateSpaceManager(privateSpaceService)

	// Fabrics Injection
	fabricClientConfig := anyctlConfigManager.NewFabricClientConfig()
	fabricClient := fabrics.NewDefaultFabricClient(*fabricClientConfig)
	fabricService := services.NewDefaultFabricService(fabricClient)
	FabricManager = managers.NewDefaultFabricManager(fabricService)

	// Scheduler Injection
	schedulerConfigClient := anyctlConfigManager.NewSchedulerClientConfig()
	schedulerClient := schedulers.NewDefaultSchedulerClient(schedulerConfigClient)
	schedulerService := services.NewDefaultSchedulerService(schedulerClient)
	SchedulerManager = managers.NewDefaultSchedulerManager(DeploymentManager, schedulerService)

	// Deployer Injection
	DeployerManager = managers.NewDefaultDeployerManager(DeploymentManager, AssetManager, OrganizationRutimeFabricManager)

	// Account Injection
	accountConfig := anyctlConfigManager.NewAccountClientConfig()
	accountClient := accounts.NewDefaultAccountClient(accountConfig)
	accountService := services.NewDefaultAccountService(accountClient)
	AccountManager = managers.NewDefaultAccountManager(accountService)
}

func exit(err error) {
	fmt.Println("ERROR: ", err.Error())
	os.Exit(1)
}
