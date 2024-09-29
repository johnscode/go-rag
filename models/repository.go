package models

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

func NewRepository(db *gorm.DB, logger *zerolog.Logger) *Repository {

	//Auto-migrate the schema
	err := db.AutoMigrate(&CorpusDoc{}, &CorpusDocChunk{}, &CorpusDocContent{}, &BM25{})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to migrate models")
	}

	return &Repository{db: db, logger: logger}
}

func (r *Repository) Close() {
	sqlDb, err := r.db.DB()
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to close database")
		return
	}
	_ = sqlDb.Close()
}

// BM25

func (r *Repository) CreateBM25(bm25 *BM25) error {
	return r.db.Create(bm25).Error
}

// CorpusDoc-related functions

func (r *Repository) CreateCorpusDoc(doc *CorpusDoc) error {
	return r.db.Create(&doc).Error
}
func (r *Repository) BulkCreateCorpusDocs(docs []CorpusDoc) error {
	return r.db.Create(&docs).Error
}

func (r *Repository) GetCorpusDocByName(docName string) (*CorpusDoc, error) {
	var device CorpusDoc
	err := r.db.Preload("Chunks").Preload("Content").Where("name = ?", docName).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *Repository) UpdateCorpusDoc(doc *CorpusDoc) error {
	return r.db.Save(doc).Error
}

func (r *Repository) DeleteCorpusDoc(docName string) error {
	return r.db.Where("name = ?", docName).Delete(&CorpusDoc{}).Error
}

//----

//func (r *Repository) CreateCorpusDocChunk(doc *CorpusDocChunk) error {
//	return r.db.Create(&doc).Error
//}
//func (r *Repository) BulkCreateCorpusDocChunks(docs []CorpusDocChunk) error {
//	return r.db.Create(&docs).Error
//}
//
//func (r *Repository) GetCorpusDocChunkByID(id uint) (*CorpusDocChunk, error) {
//	var chunk CorpusDocChunk
//	err := r.db.Where("id = ?", id).First(&chunk).Error
//	if err != nil {
//		return nil, err
//	}
//	return &chunk, nil
//}
//
//func (r *Repository) UpdateCorpusDocChunk(doc *CorpusDocChunk) error {
//	return r.db.Save(doc).Error
//}
//
//func (r *Repository) DeleteCorpusDocChunk(id uint) error {
//	return r.db.Where("id = ?", id).Delete(&CorpusDocChunk{}).Error
//}

//----

//func (r *Repository) CreateCorpusDocContent(doc *CorpusDocContent) error {
//	return r.db.Create(&doc).Error
//}
//
//func (r *Repository) GetCorpusDocContentByID(id uint) (*CorpusDocContent, error) {
//	var chunk CorpusDocContent
//	err := r.db.Where("id = ?", id).First(&chunk).Error
//	if err != nil {
//		return nil, err
//	}
//	return &chunk, nil
//}
//
//func (r *Repository) UpdateCorpusDocContent(doc *CorpusDocContent) error {
//	return r.db.Save(doc).Error
//}
//
//func (r *Repository) DeleteCorpusDocContent(id uint) error {
//	return r.db.Where("id = ?", id).Delete(&CorpusDocContent{}).Error
//}
