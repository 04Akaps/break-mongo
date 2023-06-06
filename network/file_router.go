package network

import (
	"break-mongo/document"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"sync"
)

var (
	fileApiInit     sync.Once
	fileApiInstance *fileAPI
)

type fileAPI struct {
	router *Network
	doc    *document.Document
}

func NewFileAPI(router *Network, doc *document.Document) *fileAPI {
	fileApiInit.Do(func() {
		fileApiInstance = &fileAPI{
			router: router,
			doc:    doc,
		}

		router.RegisterPOSTHandler(BASE_FILE_URI+"/upload", fileApiInstance.uploadFile)
		router.RegisterPOSTHandler(BASE_FILE_URI+"/uploadStream", fileApiInstance.uploadFileByStream)

		router.RegisterGETHandler(BASE_FILE_URI+"/retrieve", fileApiInstance.retrieveFile)
		router.RegisterGETHandler(BASE_FILE_URI+"/download", fileApiInstance.downloadFile)

		router.RegisterPUTHandler(BASE_FILE_URI+"/reName", fileApiInstance.reNameFile)

	})

	return fileApiInstance
}

func (f *fileAPI) uploadFile(c *gin.Context) {
	if file, err := os.Open("./gridFs/gridFsFile.txt"); err != nil {
		log.Println(err)
		f.router.Resp(c, 401, "File Find Failed")
	} else {
		f.doc.UploadFile(file)
	}
}

func (f *fileAPI) uploadFileByStream(c *gin.Context) {
	if file, err := os.Open("./gridFs/gridFsFile.txt"); err != nil {
		log.Println(err)
		f.router.Resp(c, 401, "File Find Failed")
	} else {
		f.doc.UploadFileByStraem(file)
	}
}

func (f *fileAPI) retrieveFile(c *gin.Context) {
	f.doc.RetrieveFile()
}

func (f *fileAPI) reNameFile(c *gin.Context) {
	f.doc.ReNameFile()
}

func (f *fileAPI) downloadFile(c *gin.Context) {
	f.doc.DownloadFile()
}
