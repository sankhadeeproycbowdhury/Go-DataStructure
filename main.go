package main

import (
	"fmt"

	"github.com/sankhadeeproycbowdhury/Go-DataStructure/queue"
)

func canReach(arr []int, start int) bool {
	visited := make(map[int]bool, len(arr))
    q := queue.Queue{}
    q.Enqueue(start)

    for q.Size() > 0 {
        currs := q.Values()
		for _, curr := range currs {
			if arr[curr] == 0 {
				return true
			}

			if curr-arr[curr] >= 0 && !visited[curr-arr[curr]] {
				q.Enqueue(curr - arr[curr])
			}

			if curr+arr[curr] < len(arr) && !visited[curr+arr[curr]] {
				q.Enqueue(curr + arr[curr])
			}

			visited[curr] = true
			q.Dequeue()
		}
	}

    return false
}

func main() {
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5)) // true
	fmt.Println(canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0)) // true
	fmt.Println(canReach([]int{3, 0, 2, 1, 2}, 2))    // false
}
