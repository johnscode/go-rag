package testhelper

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"go-rag/database"
	"gorm.io/gorm"
	"os"
	"strings"
	"testing"
	"time"
)

func SetupTestDB(t *testing.T, logger *zerolog.Logger) *gorm.DB {
	db := database.SetupTestPostgres(logger)
	//repo := models.NewRepository()

	t.Cleanup(func() {
		sqlDb, err := db.DB()
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to connect to database")
		}
		_ = sqlDb.Close()
	})

	return db
}

func ClearTestDB(t *testing.T, db *gorm.DB) {
	//Add all your models here
	err := db.Exec("TRUNCATE bm25, bm25_corpus, corpus_docs, corpus_doc_chunks, corpus_doc_contents CASCADE").Error
	if err != nil {
		t.Fatalf("Failed to clear test database: %v", err)
	}
}

func SetupLogger(ctx context.Context) *zerolog.Logger {
	var outWriter = os.Stdout
	cout := zerolog.ConsoleWriter{Out: outWriter, TimeFormat: time.RFC822}
	cout.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	// uncomment to remove timestamp from logs
	//out.FormatTimestamp = func(i interface{}) string {
	//	return ""
	//}
	baseLogger := zerolog.New(cout).With().Timestamp().Logger()
	logCtx := baseLogger.WithContext(ctx)
	l := zerolog.Ctx(logCtx)
	return l
}
