package problem0027; 

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test ( t *testing.T ) {
	var questions = []struct {
		nums []int
		target int
		ans int
	} {
			{
				[]int{ 1,2,3,4 },
				3,
				3,
			},{
				[]int{ 1,2,3,4,4,4,3,3,2 },
				2,
			7,
			},
	}

	tst := assert.New(t)
	for i, qs := range questions {
		fmt.Printf("Test %d: %v\t%d\n", i+1, qs.nums, qs.target)
		tst.Equal( qs.ans, removeElement( qs.nums, qs.target ), "  :%v", qs )
	}
}
