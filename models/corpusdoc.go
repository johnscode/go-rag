package models

type CorpusDoc struct {
	BaseModel
	Name      string           `json:"name" gorm:"unique"`
	Source    string           `json:"source"`
	ContentID uint             `json:"content_id"`
	Content   CorpusDocContent `json:"content" gorm:"foreignKey:ContentID"`
	Chunks    []CorpusDocChunk `json:"chunks" gorm:"foreignKey:DocumentID"`
}

type CorpusDocContent struct {
	BaseModel
	Content string `json:"content"`
}

type CorpusDocChunk struct {
	BaseModel
	DocumentID  uint   `json:"document_id"`
	ChunkNumber int    `json:"chunk_number"`
	Content     string `json:"content"`
}
