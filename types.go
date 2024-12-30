package storage

const (
	PathToVideo    = "video"
	PathToPhoto    = "photo"
	PathToDocument = "document"
)

const (
	TypeToVideo    = "typeVideo"
	TypeToPhoto    = "typePhoto"
	TypeToDocument = "typeDocument"
)

var (
	TypesVideo    = []string{".mp4", ".avi", ".mkv", ".webm", ".mov", ".wmv"}
	TypesPhoto    = []string{".giv", ".png", ".jpeg", ".bmp", ".jpg", ".jpe", ".jfif"}
	TypesDocument = []string{".doc", ".docx", ".pdf", "xlsx", ".xls"}
)
