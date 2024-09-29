package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"math"
	"sort"
	"strings"
)

type BM25 struct {
	BaseModel
	Corpus    []CorpusDocChunk `json:"corpus" gorm:"many2many:bm25_corpus"`
	K1        float64          `json:"k1"`
	B         float64          `json:"b"`
	AvgDocLen float64          `json:"avgDocLen"`
	// inverse all doc frequency of terms
	Idf     map[string]float64 `json:"idf" gorm:"-"`
	IdfData string             `json:"idf_data" gorm:"json"`
}

func (bm *BM25) BeforeSave(_ *gorm.DB) (err error) {
	if bm.Idf == nil {
		bm.IdfData = ""
		return nil
	}
	jsonData, err := json.Marshal(bm.Idf)
	if err != nil {
		return err
	}
	bm.IdfData = string(jsonData)
	return nil
}

func (bm *BM25) AfterFind(_ *gorm.DB) (err error) {
	if bm.IdfData == "" {
		bm.Idf = make(map[string]float64)
		return nil
	}
	err = json.Unmarshal([]byte(bm.IdfData), &bm.Idf)
	return
}

func NewBM25(corpus []CorpusDocChunk) *BM25 {
	bm25 := &BM25{
		Corpus: corpus,
		K1:     1.5,
		B:      0.75,
		Idf:    make(map[string]float64),
	}

	bm25.computeIDF()
	bm25.computeAvgDocLen()

	return bm25
}

func (bm *BM25) computeIDF() {
	numDocs := float64(len(bm.Corpus))
	docFreq := make(map[string]float64)

	for _, doc := range bm.Corpus {
		seenTerms := make(map[string]bool)
		words := strings.Fields(doc.Content)
		for _, term := range words {
			if !seenTerms[term] {
				docFreq[term]++
				seenTerms[term] = true
			}
		}
	}

	for term, count := range docFreq {
		bm.Idf[term] = math.Log((numDocs - count + 0.5) / (count + 0.5))
	}
}

func (bm *BM25) computeAvgDocLen() {
	totalLen := 0
	for _, doc := range bm.Corpus {
		totalLen += len(strings.Fields(doc.Content))
	}
	bm.AvgDocLen = float64(totalLen) / float64(len(bm.Corpus))
}

func (bm *BM25) Score(query string, doc string) float64 {
	score := 0.0
	docLen := float64(len(strings.Fields(doc)))

	termFreq := make(map[string]float64)
	for _, term := range strings.Fields(doc) {
		termFreq[term]++
	}

	for _, term := range strings.Fields(query) {
		if idf, ok := bm.Idf[term]; ok {
			tf := termFreq[term]
			numerator := tf * (bm.K1 + 1)
			denominator := tf + bm.K1*(1-bm.B+bm.B*docLen/bm.AvgDocLen)
			score += idf * numerator / denominator
		}
	}

	return score
}

func (bm *BM25) RankDocuments(query string) []int {
	scores := make([]float64, len(bm.Corpus))
	for i, doc := range bm.Corpus {
		scores[i] = bm.Score(query, doc.Content)
	}

	// Create a slice of indices
	indices := make([]int, len(scores))
	for i := range indices {
		indices[i] = i
	}

	// Sort indices based on scores (descending order)
	sort.Slice(indices, func(i, j int) bool {
		return scores[indices[i]] > scores[indices[j]]
	})

	return indices
}
