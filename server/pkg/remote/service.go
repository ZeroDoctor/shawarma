package remote

import "github.com/zerodoctor/shawarma/internal/db"

var remoteServicesMap map[string]GitRemoteDriver = make(map[string]GitRemoteDriver)
var remoteServices []string

func Register(remoteName string, remote GitRemoteDriver) {
	remoteServicesMap[remoteName] = remote
	remoteServices = append(remoteServices, remoteName)
}

func Setup(db db.DB) {
	for _, remote := range remoteServicesMap {
		remote.Setup(db)
	}
}

func GetRemoteService(name string) GitRemoteDriver {
	return remoteServicesMap[name]
}

func GetRemoteNames() []string {
	return remoteServices
}
