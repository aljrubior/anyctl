package cmd

import (
	"fmt"
	"github.com/aljrubior/anyctl/conf"
	"github.com/aljrubior/anyctl/managers"
	"github.com/aljrubior/anyctl/utils"
	"github.com/aljrubior/anyctl/wires"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var Console utils.Console

var AnyctlConfig *conf.AnyctlConfig

var ConfigManager managers.ConfigManager
var AssetManager managers.AssetManager
var DeploymentManager managers.DeploymentManager
var TargetManager managers.TargetManager
var OrganizationRutimeFabricManager managers.OrganizationRuntimeFabricManager
var OrganizationPrivateSpaceManager managers.OrganizationPrivateSpaceManager
var SharedSpaceManager managers.SharedSpaceManager
var PrivateSpaceManager managers.PrivateSpaceManager
var FabricManager managers.FabricManager
var DeployerManager managers.DeployerManager
var SchedulerManager managers.SchedulerManager
var AccountManager managers.AccountManager

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
	anyctlConfigManager, err := wires.InitializeAnyctlConfigManager()

	if err != nil {
		exit(err)
	}

	// Config Injection
	ConfigManager = wires.InitializeConfigManager()

	// Asset Injection
	AssetManager, err = wires.InitializeAssetManager(anyctlConfigManager.AssetClientConfig)

	if err != nil {
		exit(err)
	}

	// Deployment Injection
	DeploymentManager, err = wires.InitializeDeploymentManager(anyctlConfigManager.DeploymentClientConfig, AssetManager)

	if err != nil {
		exit(err)
	}

	// Target Injection
	TargetManager, err = wires.InitializeTargetManager(anyctlConfigManager.TargetClientConfig)

	if err != nil {
		exit(err)
	}

	// Organization Runtime Fabric Injection
	OrganizationRutimeFabricManager, err = wires.InitializeOrganizationRuntimeFabricManager(anyctlConfigManager.OrganizationRuntimeFabricConfig)

	if err != nil {
		exit(err)
	}

	// Organization Private Space Injection
	OrganizationPrivateSpaceManager, err = wires.InitializeOrganizationPrivateSpaceManager(anyctlConfigManager.OrganizationRuntimeFabricConfig)

	if err != nil {
		exit(err)
	}

	// Shared Space Injection
	SharedSpaceManager, err = wires.InitializeSharedSpaceManager(anyctlConfigManager.GetSharedSpaceClientConfig())

	if err != nil {
		exit(err)
	}

	// Private Space Injection
	PrivateSpaceManager, err = wires.InitializePrivateSpaceManager(anyctlConfigManager.GetPrivateSpaceClientConfig())

	if err != nil {
		exit(err)
	}

	// Fabrics Injection
	FabricManager, err = wires.InitializeFabricManager(anyctlConfigManager.GetFabricClientConfig())

	if err != nil {
		exit(err)
	}

	// Scheduler Injection
	SchedulerManager, err = wires.InitializeSchedulerManager(anyctlConfigManager.GetSchedulerClientConfig(), DeploymentManager)

	if err != nil {
		exit(err)
	}

	// Deployer Injection
	DeployerManager, err = wires.InitializeDeployerManager(DeploymentManager, AssetManager, OrganizationRutimeFabricManager)

	if err != nil {
		exit(err)
	}

	// Account Injection
	AccountManager, err = wires.InitializeAccountManager(anyctlConfigManager.GetAccountClientConfig())

	if err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println("Error: ", err.Error())
	os.Exit(1)
}
