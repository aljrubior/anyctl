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
var OrganizationRuntimeFabricManager managers.OrganizationRuntimeFabricManager
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
	if AssetManager, err = wires.InitializeAssetManager(anyctlConfigManager.GetAssetClientConfig()); err != nil {
		exit(err)
	}

	// Deployment Injection
	if DeploymentManager, err = wires.InitializeDeploymentManager(anyctlConfigManager.GetDeploymentConfigClient(), AssetManager); err != nil {
		exit(err)
	}

	// Target Injection
	if TargetManager, err = wires.InitializeTargetManager(anyctlConfigManager.GetTargetClientConfig()); err != nil {
		exit(err)
	}

	// Organization Runtime Fabric Injection
	if OrganizationRuntimeFabricManager, err = wires.InitializeOrganizationRuntimeFabricManager(anyctlConfigManager.GetOrganizationRuntimeFabricClientConfig()); err != nil {
		exit(err)
	}

	// Organization Private Space Injection
	if OrganizationPrivateSpaceManager, err = wires.InitializeOrganizationPrivateSpaceManager(anyctlConfigManager.GetOrganizationRuntimeFabricClientConfig()); err != nil {
		exit(err)
	}

	// Shared Space Injection
	if SharedSpaceManager, err = wires.InitializeSharedSpaceManager(anyctlConfigManager.GetSharedSpaceClientConfig()); err != nil {
		exit(err)
	}

	// Private Space Injection
	if PrivateSpaceManager, err = wires.InitializePrivateSpaceManager(anyctlConfigManager.GetPrivateSpaceClientConfig()); err != nil {
		exit(err)
	}

	// Fabrics Injection
	if FabricManager, err = wires.InitializeFabricManager(anyctlConfigManager.GetFabricClientConfig()); err != nil {
		exit(err)
	}

	// Scheduler Injection
	if SchedulerManager, err = wires.InitializeSchedulerManager(anyctlConfigManager.GetSchedulerClientConfig(), DeploymentManager); err != nil {
		exit(err)
	}

	// Deployer Injection
	if DeployerManager, err = wires.InitializeDeployerManager(DeploymentManager, AssetManager, OrganizationRuntimeFabricManager); err != nil {
		exit(err)
	}

	// Account Injection
	if AccountManager, err = wires.InitializeAccountManager(anyctlConfigManager.GetAccountClientConfig()); err != nil {
		exit(err)
	}

}

func exit(err error) {
	fmt.Println("Error: ", err.Error())
	os.Exit(1)
}
