package mincostclimbingstairs

// https://leetcode.com/problems/min-cost-climbing-stairs/
// Difficulty: Easy

func minCostClimbingStairs(cost []int) int {
	prev := cost[1]
	prevprev := cost[0]
	for i := 2; i < len(cost); i++ {
		curr := cost[i] + min(prev, prevprev)
		prevprev = prev
		prev = curr
	}
	return min(prev, prevprev)
}
