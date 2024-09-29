package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_CreateCorpusDoc(t *testing.T) {
	_, _, repo := setup(t)

	cd := quickFoxDoc()
	err := repo.CreateCorpusDoc(&cd)
	assert.NoError(t, err)
	assert.NotZero(t, cd.ID)
}

func TestRepository_GetCorpusDoc(t *testing.T) {
	_, _, repo := setup(t)

	cd := quickFoxDoc()
	_ = repo.CreateCorpusDoc(&cd)
	doc, err := repo.GetCorpusDocByName(cd.Name)
	assert.NoError(t, err)
	assert.Equal(t, cd, *doc)
	assert.Equal(t, 1, len(doc.Chunks))
	assert.Equal(t, quickFoxStr, doc.Content.Content)
}
