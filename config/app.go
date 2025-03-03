package config

import "os"


var serverEnvName string = os.Getenv("SERVER_ENV_NAME")
const LogDir= "/go-ecommerce/log"
const AppCodeName = "ecommerce"



const (
	ServerTypeNameLocal = "local"
	ServerTypeNameDevelop = "develop"
	ServerTypeNameClosedAlpha = "closed_alpha"
	ServerTypeNameStaging = "staging"
	ServerTypeNameStagingMgr = "staging-mgr"
	ServerTypeNameRelease = "release"
	ServerTypeNameReleaseMgr = "release-mgr"
	
)







func GetServerEnvName() string {
	return serverEnvName
}

func GetIsLocal() bool {
	switch GetServerEnvName() {
	case ServerTypeNameLocal:
		return true	
	}
	return false
}


func GetIsDevelop() bool {
	switch GetServerEnvName() {
	case ServerTypeNameDevelop:
		return true
	}
	return false
}
func GetIsClosedAlpha() bool {
	return GetServerEnvName() == ServerTypeNameClosedAlpha
}



func GetIsDebugMode()  bool {
	name := GetServerEnvName()
	switch name {
	case ServerTypeNameRelease, ServerTypeNameReleaseMgr:
		return false
	}
	return true
}

func RouteManagementPage() bool{
	switch GetServerEnvName() {
	case ServerTypeNameStaging, ServerTypeNameRelease:
		return false
	}
	return true
}

func GetIsManagement() bool {
	switch GetServerEnvName() {
	case ServerTypeNameStagingMgr, ServerTypeNameReleaseMgr:
		return true
	}
	return false
}
func GetServerRootUrl() string {
	return ""
}


func GetServerApiRootUrl() string {
	switch GetServerEnvName() {
	case ServerTypeNameLocal:
		return "http://localhost:8082/v1/2024"
	}
	return ""
}


func getBlockchainApiURL() string {
	switch GetServerEnvName() {
	case ServerTypeNameLocal:
		return "http://194.233.82.172:3001/"
	case ServerTypeNameClosedAlpha:
		return "http://194.233.82.172:3001/"
	case ServerTypeNameDevelop:
		return "http://194.233.82.172:3001/"
	default:
		return "http://194.233.82.172:3001/"
	}
}

func GetAppUrl() string {
	switch GetServerEnvName() {
	case ServerTypeNameLocal:
		return "http://localhost:8082/"
	}
	return ""
}


