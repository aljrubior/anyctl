package utils

import "fmt"

type Console struct {
}

func (this Console) LogError(err error) {
	println(fmt.Sprintf("ERROR: %s", err.Error()))
}

func (this Console) LogAccessTokenUpdate() {
	println("Access token updated.")
}

func (this Console) LogRunSchedulerSuccess(schedulerName string) {
	println(fmt.Sprintf("Scheduler '%s' triggered.", schedulerName))
}

func (this Console) LogInvalidParameters() {
	println("ERROR: Invalid parameters.")
}

func (this Console) LogUpdatedCurrentEnvironmentSuccess(environmentName, filePath string) {
	println(fmt.Sprintf("Updated current environment '%s' in '%s'.", environmentName, filePath))
}
