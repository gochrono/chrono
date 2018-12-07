package chronolib

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/jinzhu/now"
    "time"
    "encoding/hex"
	"testing"
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
    uuid, _ := hex.DecodeString(h)
    return Frame{uuid, project, time.Time{}, time.Time{}, []string{}, []string{}}
}

func TestGetFrameByShortHex(t *testing.T) {
    frames := []Frame{
        createFrame("cf23df2207d99a74fbe169e3eba035e633b65d94", "something"),
        createFrame("ddb9abade9487ee132588da9c2479894c8c0e208", "timevault"),
        createFrame("c394b77210dfd79161b9492d224541258ee3a9c0", "development"),
    }
    
    index, frame, err := GetFrameByShortHex(frames, "cf23df")
    assert.Equal(t, 0, index)
    assert.Equal(t, "something", frame.Project)
    assert.Equal(t, err, nil)
    
    index, frame, err = GetFrameByShortHex(frames, "ddb9ab")
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