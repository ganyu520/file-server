package meta

// 文件元信息
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

// 通过FileSha1查找文件元信息
var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

// GetFileMeta:通过FileSha1获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
