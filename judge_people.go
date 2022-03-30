/*In a town, there are n people labeled from 1 to n. There is a rumor that one of these people is secretly the town judge.
If the town judge exists, then:
The town judge trusts nobody.Everybody (except for the town judge) trusts the town judge.There is exactly one person that satisfies properties 1 and 2.
You are given an array trust where trust[i] = [ai, bi] representing that the person labeled ai trusts the person labeled bi.
Return the label of the town judge if the town judge exists and can be identified, or return -1 otherwise.


Example 1:
Input: n = 2, trust = [[1,2]]
Output: 2

Example 2:
Input: n = 3, trust = [[1,3],[2,3]]
Output: 3

Example 3:
Input: n = 3, trust = [[1,3],[2,3],[3,1]]
Output: -1*/

package main

import (
	"fmt"
)

func find_judge(n int, trust [][]int) int {
	people := make(map[int]int)
	judge := make(map[int]int)

	for _, val := range trust {
		people[val[0]]++
		judge[val[1]]++
	}

	res := -1
	for i, v := range judge {
		if v == n-1 && people[i] == 0 {
			res = i
			break
		}
	}

	return res
}

func main() {
	fmt.Println(find_judge(2, [][]int{{1, 2}}))
	fmt.Println(find_judge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
}
