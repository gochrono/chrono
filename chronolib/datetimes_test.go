package chronolib

import (
	"github.com/gofrs/uuid"
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

func TestSubFrameForTimespan(t *testing.T) {
	start1, end1 := GetTimespanForDay("2019-12-20")
	start2, end2 := GetTimespanForDay("2019-12-21")

	frameStart, _ := now.Parse("2019-12-20 22:00")
	frameEnd, _ := now.Parse("2019-12-21 1:00")

	frame := Frame{
		UUID: uuid.Must(uuid.NewV4()).String(),
		Project: "test",
		StartedAt: frameStart,
		EndedAt: frameEnd,
		UpdatedAt: frameEnd,
		Tags: nil,
		Notes: nil,
	}

	subFrame1 := SubFrameForTimespan(frame, start1, end1)
	assert.Equal(t, subFrame1.StartedAt, frame.StartedAt)
	assert.Equal(t, subFrame1.EndedAt, end1)

	subFrame2 := SubFrameForTimespan(frame, start2, end2)
	assert.Equal(t, subFrame2.StartedAt, start2)
	assert.Equal(t, subFrame2.EndedAt, frame.EndedAt)
}
