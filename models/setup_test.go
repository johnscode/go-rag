package models

import (
	"context"
	"github.com/rs/zerolog"
	"go-rag/database"
	"go-rag/testhelper"
	"gorm.io/gorm"
	"testing"
)

func setup(t *testing.T) (*zerolog.Logger, *gorm.DB, *Repository) {
	logger := testhelper.SetupLogger(context.Background())
	db := database.SetupTestPostgres(logger)
	repo := NewRepository(db, logger)
	testhelper.ClearTestDB(t, db)
	return logger, db, repo
}

const quickFoxStr = "the quick brown fox jumps over the lazy dog"

var quickFoxStrs = []string{"the quick brown fox jumps", " over the lazy dog"}

func quickFoxDoc() CorpusDoc {

	return CorpusDoc{
		BaseModel: BaseModel{},
		Name:      "testdoc",
		Source:    "localhost",
		Content: CorpusDocContent{
			BaseModel: BaseModel{},
			Content:   quickFoxStr,
		},
		Chunks: []CorpusDocChunk{
			{
				BaseModel:   BaseModel{},
				ChunkNumber: 0,
				Content:     quickFoxStr,
			},
		},
	}
}

func FourChunks() CorpusDoc {
	var chunks = []CorpusDocChunk{
		{
			BaseModel:   BaseModel{},
			ChunkNumber: 0,
			Content:     "one",
		},
		{
			BaseModel:   BaseModel{},
			ChunkNumber: 1,
			Content:     " two",
		},
		{
			BaseModel:   BaseModel{},
			ChunkNumber: 1,
			Content:     " three",
		},
		{
			BaseModel:   BaseModel{},
			ChunkNumber: 1,
			Content:     " four",
		},
	}
	return CorpusDoc{
		BaseModel: BaseModel{},
		Name:      "testdoc",
		Source:    "localhost",
		Content: CorpusDocContent{
			BaseModel: BaseModel{},
			Content:   "one two three four",
		},
		Chunks: chunks,
	}
}
