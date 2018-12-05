package chronolib

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeTags(t *testing.T) {
	input := []string{"tag", "tag Normal", "TAG-x"}
	expected := []string{"tag", "tag-normal", "tag-x"}

	output := NormalizeTags(input)

	for i := range output {
		assert.Equal(t, output[i], expected[i])
	}
}

func TestIsTags(t *testing.T) {
	good := []string{"+tag", "+tag Normal", "+TAG-x"}
	bad := []string{"+tag", "tag Normal", "+TAG-x"}

	goodErr := CheckTags(good)
	badErr := CheckTags(bad)
	assert.Equal(t, goodErr, nil)
	assert.Equal(t, badErr, errors.New("invalid tag: tag Normal"))
}
