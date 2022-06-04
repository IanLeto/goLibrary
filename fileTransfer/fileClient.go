package fileTransfer

import (
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
)

func TransferFile(localPath, fileName , remotePath, remoteFile, url string)  {
	clientConfig ,_ := auth
}