package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicCases(t *testing.T) {
	input := NewBasicRepoInput{
		modelType:        "basic",
		readModelBuilder: ReadModelBuilder,
		fetcher:          DefaultDBFetcher,
	}

	repo := NewBasicRepo(input)

	agg := repo.GetAggregate("i")
	assert.Equal(t, "1", agg.ID)
}
