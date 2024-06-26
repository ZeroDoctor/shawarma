package remote

import "github.com/zerodoctor/shawarma/internal/db"

var remoteServices map[string]GitRemote

func Register(remoteName string, remote GitRemote) {
	remoteServices[remoteName] = remote
}

func Setup(db db.DB) {
	for _, remote := range remoteServices {
		remote.Setup(db)
	}
}

func GetRemoteService(name string) GitRemote {
	return remoteServices[name]
}
