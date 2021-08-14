package entity

// FileMeta : 文件元信息结构
type FileMeta struct {
	FileSha1 string `xorm:"file_sha1"`
	FileName string `xorm:"file_name"`
	FileSize int64  `xorm:"file_size"`
	Location string `xorm:"location"`
	UploadAt string `xorm:"upload_at"`
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta : 新增/更新文件元信息
func (f FileMeta) UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta : 通过sha1值获取文件的元信息对象
func (f FileMeta) GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
