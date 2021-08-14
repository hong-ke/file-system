package service

import (
	"filesystem/entity"
	"filesystem/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"time"
)

type UploadService struct {
	uploadPath string
}

func NewUploadService(uploadPath string) (*UploadService, error) {
	if len(uploadPath) <= 0 {
		return nil, errors.New("upload path 不能为空")
	}
	return &UploadService{uploadPath: uploadPath}, nil
}

func (u *UploadService) SaveFile(file multipart.File, fileName string) error {
	uploadPath := path.Join(u.uploadPath, fileName)
	newFile, err := os.Create(uploadPath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	defer newFile.Close()
	fileMeta := &entity.FileMeta{
		FileName: fileName,
		Location: uploadPath,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		return errors.WithStack(err)
	}
	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)
	fileMeta.UpdateFileMeta(*fileMeta)
	log.Infof("meta: %s", fileMeta.FileSha1)
	return nil
}

func (u *UploadService) GetFileMeta(fileSha1 string) (*entity.FileMeta, error) {
	meta := entity.FileMeta{}
	meta = meta.GetFileMeta(fileSha1)
	return &meta, nil
}

func (u *UploadService) DownloadFile(fileSha1 string) ([]byte, *string, error) {
	meta, err := u.GetFileMeta(fileSha1)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	file, err := os.Open(meta.Location)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return data, &meta.FileName, nil
}
