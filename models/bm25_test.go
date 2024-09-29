package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CreateBM25(t *testing.T) {
	t.Run("single document", func(t *testing.T) {
		_, _, repo := setup(t)

		cd := quickFoxDoc()
		err := repo.CreateCorpusDoc(&cd)
		assert.NoError(t, err)
		assert.NotZero(t, cd.ID)

		bm := NewBM25(cd.Chunks)
		err = repo.CreateBM25(bm)
		assert.NoError(t, err)
		assert.NotZero(t, cd.ID)
	})

	t.Run("single document, 2 chunks", func(t *testing.T) {
		logger, _, repo := setup(t)
		cd := FourChunks()
		err := repo.CreateCorpusDoc(&cd)

		bm := NewBM25(cd.Chunks)
		err = repo.CreateBM25(bm)
		assert.NoError(t, err)
		assert.NotZero(t, cd.ID)
		assert.Zero(t, bm.Idf["quick"])

		rank := bm.RankDocuments("one")
		assert.Equal(t, 0, rank[0])
		logger.Info().Msg(fmt.Sprintf("rank: %+v", rank))
		rank = bm.RankDocuments("two")
		assert.Equal(t, 1, rank[0])
		logger.Info().Msg(fmt.Sprintf("rank: %+v", rank))
		rank = bm.RankDocuments("three")
		logger.Info().Msg(fmt.Sprintf("rank: %+v", rank))
		assert.Equal(t, 2, rank[0])
	})
}
