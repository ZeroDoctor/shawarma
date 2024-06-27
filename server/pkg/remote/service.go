package remote

import "github.com/zerodoctor/shawarma/internal/db"

var remoteServicesMap map[string]GitRemote
var remoteServices []string

func Register(remoteName string, remote GitRemote) {
	remoteServicesMap[remoteName] = remote
	remoteServices = append(remoteServices, remoteName)
}

func Setup(db db.DB) {
	for _, remote := range remoteServicesMap {
		remote.Setup(db)
	}
}

func GetRemoteService(name string) GitRemote {
	return remoteServicesMap[name]
}

func GetRemoteNames() []string {
	return remoteServices
}
