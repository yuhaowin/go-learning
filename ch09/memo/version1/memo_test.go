package memo_test

import (
	"testing"

	"github.com/yuhaowin/go-learning/ch09/memo/memotest"
	memo "github.com/yuhaowin/go-learning/ch09/memo/version1"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
