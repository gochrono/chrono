package chronolib

import (
	"github.com/jinzhu/now"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestContainsMorethanOneBooleanFlag(t *testing.T) {
	assert.Equal(t, false, ContainsMoreThanOneBooleanFlag(false))
	assert.Equal(t, false, ContainsMoreThanOneBooleanFlag(false, false))
	assert.Equal(t, false, ContainsMoreThanOneBooleanFlag(true))
	assert.Equal(t, false, ContainsMoreThanOneBooleanFlag(false, true, false))
	assert.Equal(t, true, ContainsMoreThanOneBooleanFlag(false, true, true))
	assert.Equal(t, true, ContainsMoreThanOneBooleanFlag(true, false, true, true))
}

func createFrame(h string, project string) Frame {
	return Frame{h, project, time.Time{}, time.Time{}, time.Time{}, []string{}, []string{}}
}

func createFrameWithTime(h string, project string, startTime string) Frame {
	fStart, _ := now.Parse(startTime)
	return Frame{h, project, fStart, time.Time{}, fStart, []string{}, []string{}}
}

func TestGetFrameByShortHex(t *testing.T) {
	frames := []Frame{
		createFrame("cf23df2207d99a74fbe169e3eba035e633b65d94", "something"),
		createFrame("ddb9abade9487ee132588da9c2479894c8c0e208", "timevault"),
		createFrame("c394b77210dfd79161b9492d224541258ee3a9c0", "development"),
	}

	index, frame, err := GetFrameByShortHex(frames, "cf23df2")
	assert.Equal(t, 0, index)
	assert.Equal(t, "something", frame.Project)
	assert.Equal(t, err, nil)

	index, frame, err = GetFrameByShortHex(frames, "ddb9aba")
	assert.Equal(t, 1, index)
	assert.Equal(t, "timevault", frame.Project)
	assert.Equal(t, err, nil)
}

func TestGetFrameByShortHexMissingFrame(t *testing.T) {
	frames := []Frame{
		createFrame("cf23df2207d99a74fbe169e3eba035e633b65d94", "something"),
		createFrame("ddb9abade9487ee132588da9c2479894c8c0e208", "timevault"),
		createFrame("c394b77210dfd79161b9492d224541258ee3a9c0", "development"),
	}

	_, _, err := GetFrameByShortHex(frames, "asdasd")
	assert.NotEqual(t, err, nil)
}

func TestSortFramesByDate(t *testing.T) {
	frames := []Frame{
		createFrameWithTime("cf23df2207d99a74fbe169e3eba035e633b65d94", "something", "10:25"),
		createFrameWithTime("ddb9abade9487ee132588da9c2479894c8c0e208", "timevault", "10:10"),
		createFrameWithTime("c394b77210dfd79161b9492d224541258ee3a9c0", "development", "10:05"),
	}

	SortFramesByDate(frames)
	assert.Equal(t, "development", frames[0].Project)
	assert.Equal(t, "timevault", frames[1].Project)
	assert.Equal(t, "something", frames[2].Project)
}

func TestSortTimeMapKeys(t *testing.T) {
	firstTime, _ := now.Parse("2018-9-25")
	firstFrames := []Frame{
		createFrameWithTime("cf23df2207d99a74fbe169e3eba035e633b65d94", "timevault", "2018-9-23 12:50"),
	}
	secondTime, _ := now.Parse("2018-9-23")
	secondFrames := []Frame{
		createFrameWithTime("cf23df2207d99a74fbe169e3eba035e633b65d94", "something", "2018-9-24 12:50"),
	}
	timemap := map[time.Time][]Frame{
		firstTime:  firstFrames,
		secondTime: secondFrames,
	}

	keys := SortTimeMapKeys(&timemap)
	firstExpectedTime, _ := now.Parse("2018-9-23")
	secondExpectedTime, _ := now.Parse("2018-9-25")
	assert.Equal(t, firstExpectedTime, keys[0])
	assert.Equal(t, secondExpectedTime, keys[1])
}
