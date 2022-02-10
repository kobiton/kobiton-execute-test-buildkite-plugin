package model

import (
	"os"
	"strconv"
)

type StepConfig struct {
	kobiUsername          string
	kobiApiKey            string
	executorUrl           string
	executorUsername      string
	executorPassword      string
	gitRepoUrl            string
	gitRepoBranch         string
	gitSSHKey             string
	kobiAppId             string
	useCustomDevice       bool
	deviceName            string
	devicePlatformVersion string
	devicePlatformName    string
	rootDirectory         string
	commands              string
	waitForExecution      bool
	logType               string
}

func (stepConfig *StepConfig) Init() {

	stepConfig.kobiUsername = os.Getenv("KOBI_USERNAME")
	stepConfig.kobiApiKey = os.Getenv("KOBI_API_KEY")
	stepConfig.executorUrl = os.Getenv("EXECUTOR_URL")
	stepConfig.executorUsername = os.Getenv("EXECUTOR_USERNAME")
	stepConfig.executorPassword = os.Getenv("EXECUTOR_PASSWORD")
	stepConfig.gitRepoUrl = os.Getenv("GIT_REPO_URL")
	stepConfig.gitRepoBranch = os.Getenv("GIT_REPO_BRANCH")
	stepConfig.gitSSHKey = os.Getenv("GIT_REPO_SSH_KEY")
	stepConfig.kobiAppId = os.Getenv("APP_ID")
	stepConfig.useCustomDevice, _ = strconv.ParseBool(os.Getenv("USE_CUSTOM_DEVICE"))
	stepConfig.deviceName = os.Getenv("DEVICE_NAME")
	stepConfig.devicePlatformVersion = os.Getenv("DEVICE_PLATFORM_VERSION")
	stepConfig.devicePlatformName = os.Getenv("DEVICE_PLATFORM")
	stepConfig.rootDirectory = os.Getenv("ROOT_DIRECTORY")
	stepConfig.commands = os.Getenv("COMMAND")
	stepConfig.waitForExecution, _ = strconv.ParseBool(os.Getenv("WAIT_FOR_EXECUTION"))

	switch os.Getenv("LOG_TYPE") {
	case "output":
		stepConfig.logType = "out"
	case "error":
		stepConfig.logType = "error"
	default:
		stepConfig.logType = "all"
	}
}

func (stepConfig *StepConfig) GetKobiUsername() string {
	return stepConfig.kobiUsername
}

func (stepConfig *StepConfig) GetKobiPassword() string {
	return stepConfig.kobiApiKey
}

func (stepConfig *StepConfig) GetExecutorUrl() string {
	return stepConfig.executorUrl
}

func (stepConfig *StepConfig) GetExecutorUsername() string {
	return stepConfig.executorUsername
}

func (stepConfig *StepConfig) GetExecutorPassword() string {
	return stepConfig.executorPassword
}

func (stepConfig *StepConfig) GetGitRepoUrl() string {
	return stepConfig.gitRepoUrl
}

func (stepConfig *StepConfig) GetGitRepoBranch() string {
	return stepConfig.gitRepoBranch
}

func (stepConfig *StepConfig) GetGitSSHKey() string {
	return stepConfig.gitSSHKey
}

func (stepConfig *StepConfig) GetKobiAppId() string {
	return stepConfig.kobiAppId
}

func (stepConfig *StepConfig) IsUseCustomDevices() bool {
	return stepConfig.useCustomDevice
}

func (stepConfig *StepConfig) GetDeviceName() string {
	return stepConfig.deviceName
}

func (stepConfig *StepConfig) GetDevicePlatformVersion() string {
	return stepConfig.devicePlatformVersion
}

func (stepConfig *StepConfig) GetDevicePlatformname() string {
	return stepConfig.devicePlatformName
}

func (stepConfig *StepConfig) GetRootDirectory() string {
	return stepConfig.rootDirectory
}

func (stepConfig *StepConfig) GetCommands() string {
	return stepConfig.commands
}

func (stepConfig *StepConfig) IsWaitForExecution() bool {
	return stepConfig.waitForExecution
}

func (stepConfig *StepConfig) GetLogType() string {
	return stepConfig.logType
}
