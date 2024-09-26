package GinX

import (
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func FormMultiFiles(c *gin.Context, key string) (bs [][]byte, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	files := form.File[key]
	if len(files) == 0 {
		return
	}

	var oFile multipart.File
	var b []byte
	for _, file := range files {
		oFile, err = file.Open()
		if err != nil {
			return
		}
		defer oFile.Close()
		b, err = io.ReadAll(oFile)
		if err != nil {
			return
		}

		bs = append(bs, b)
	}
	return
}

func FormMultiFilesOpenAll(c *gin.Context, key string) (ofs []multipart.File, err error) {
	form, err := c.MultipartForm()
	if err != nil {
		return
	}
	files := form.File[key]
	if len(files) == 0 {
		return
	}

	var oFile multipart.File
	for _, file := range files {
		oFile, err = file.Open()
		if err != nil {
			return
		}
		ofs = append(ofs, oFile)
	}
	return
}

func CloseAllFiles(files []multipart.File) {
	for _, file := range files {
		file.Close()
	}
}
