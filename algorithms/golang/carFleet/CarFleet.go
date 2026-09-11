// Source : https://leetcode.com/problems/car-fleet
// Author : BradleyZhang
// Date   : 2026-09-10

/*****************************************************************************************************
 *
 * There are n cars at given miles away from the starting mile 0, traveling to reach the mile target.
 * You are given two integer arrays position and speed, both of length n, where position[i] is the
 * starting mile of the i^th car and speed[i] is the speed of the i^th car in miles per hour.
 * A car cannot pass another car, but it can catch up and then travel next to it at the speed of the
 * slower car.
 * A car fleet is a single car or a group of cars driving next to each other. The speed of the car
 * fleet is the minimum speed of any car in the fleet.
 * If a car catches up to a car fleet at the mile target, it will still be considered as part of the
 * car fleet.
 * Return the number of car fleets that will arrive at the destination.
 *
 * Example 1:
 * Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
 * Output: 3
 * Explanation:
 * 	The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12.
 * The fleet forms at target.
 * 	The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by
 * itself.
 * 	The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6.
 * The fleet moves at speed 1 until it reaches target.
 *
 * Example 2:
 * Input: target = 10, position = [3], speed = [3]
 * Output: 1
 * Explanation:
 * There is only one car, hence there is only one fleet.
 *
 * Example 3:
 * Input: target = 100, position = [0,2,4], speed = [4,2,1]
 * Output: 1
 * Explanation:
 * 	The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4.
 * The car starting at 4 (speed 1) travels to 5.
 * 	Then, the fleet at 4 (speed 2) and the car at position 5 (speed 1) become one fleet,
 * meeting each other at 6. The fleet moves at speed 1 until it reaches target.
 *
 * Constraints:
 * 	n == position.length == speed.length
 * 	1 <= n <= 10^5
 * 	0 < target <= 10^6
 * 	0 <= position[i] < target
 * 	All the values of position are unique.
 * 	0 < speed[i] <= 10^6
 ******************************************************************************************************/
package carfleet

import "sort"

// carI只要target前能追上前面车队，就加入,若追不上自己就是新车队
func carFleet(target int, position []int, speed []int) int {
	cars := make([]car, len(position))
	for i := range position {
		cars[i] = car{
			position: position[i],
			speed:    speed[i],
		}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})
	result := 1
	fleet := cars[0]
	for _, c := range cars[1:] {
		if !canJoinFleet(fleet, c, target) {
			fleet = c
			result++
		}
	}
	return result
}

type car struct {
	position int
	speed    int
}

func canJoinFleet(fleet, curCar car, target int) bool {
	hourToTargetFront := float64(target-fleet.position) / float64(fleet.speed)
	hourToTargetBehind := float64(target-curCar.position) / float64(curCar.speed)
	return hourToTargetBehind <= hourToTargetFront
}

// 其他写法
func carFleet2(target int, position []int, speed []int) int {
	type car struct {
		position int
		speed    int
	}

	cars := make([]car, len(position))
	for i := range position {
		cars[i] = car{
			position: position[i],
			speed:    speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].position > cars[j].position
	})

	fleets := 0
	lastArrivalTime := 0.0

	for _, car := range cars {
		arrivalTime := float64(target-car.position) / float64(car.speed)

		if arrivalTime > lastArrivalTime {
			fleets++
			lastArrivalTime = arrivalTime
		}
	}

	return fleets
}
