package filesystem

import (
	"mime/multipart"
)

type FileDriver interface {
	PutFile(file *multipart.FileHeader, dirPath string, filename string) (string, error)
}
