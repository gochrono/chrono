package chronolib

import (
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInsideTimeRange(t *testing.T) {
	start, _ := now.Parse("2018-10-11")
	end, _ := now.Parse("2018-10-9")

	point, _ := now.Parse("2018-10-10")

	result := IsTimeInTimespan(point, start, end)
	assert.Equal(t, result, true)
}
