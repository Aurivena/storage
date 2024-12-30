package storage

const (
	PathToVideo    = "video"
	PathToPhoto    = "photo"
	PathToDocument = "document"
	PathToArchive  = "archive"
)

const (
	TypeToVideo    = "typeVideo"
	TypeToPhoto    = "typePhoto"
	TypeToDocument = "typeDocument"
	TypeToArchive  = "typeArchive"
)

var (
	TypesVideo    = []string{".mp4", ".avi", ".mkv", ".webm", ".mov", ".wmv"}
	TypesPhoto    = []string{".giv", ".png", ".jpeg", ".bmp", ".jpg", ".jpe", ".jfif"}
	TypesDocument = []string{".doc", ".docx", ".pdf", "xlsx", ".xls"}
	TypesArchive  = []string{".zip", ".rar", ".7z", ".tar", ".gz", ".xz", ".bzw"}
)
