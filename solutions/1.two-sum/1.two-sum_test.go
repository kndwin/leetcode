package problem0001; 

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test ( t *testing.T ) {
	var questions = []struct {
		nums []int
		target int
		ans []int
	} {
			{
				[]int{ 1,2,3,4 },
				7,
				[]int{ 2,3 },
			},{
				[]int{ 1,2,3,4 },
				6,
				[]int{ 1,3 }, 
			},{
				[]int{ -1,-2,3,4 },
				1,
				[]int{ 1,2 }, 
			},
	}

	tst := assert.New(t)
	for i, qs := range questions {
		fmt.Printf("Test %d: %v\t%d\n", i+1, qs.nums, qs.target)
		tst.Equal( qs.ans, twoSum( qs.nums, qs.target ), "  :%v", qs )
	}
}
