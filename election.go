package main

import (
	"fmt"
)

//Pair storing height and time data
type Pair struct {
	height int
	time   int
}

//List define list for storing pair
type IotexList struct {
	tableHeighs map[int]bool
	tableTimes  map[int]bool
	heights     []int
	times       []int
}

func (list *IotexList) checkContainHeight(height int) bool {
	return list.tableHeighs[height] == true
}

func (list *IotexList) checkContainTime(time int) bool {
	return list.tableHeighs[time] == true
}

//GetList return list
func (list *IotexList) GetList() []Pair {
	if list.isConflict() {
		return []Pair{}
	}
	pairs := []Pair{}

	for i, height := range list.heights {
		pairs = append(pairs, Pair{height, list.times[i]})
	}
	return pairs
}

func (list *IotexList) printPair() {
	fmt.Println(list.heights)
	fmt.Println(list.times)
}

func (list *IotexList) nearestHeightBefore(time int) int {
	index := list.binarySearch(list.times, time-1, 0, len(list.times)-1)
	if index == -1 {
		if list.times[0] < time {
			return list.heights[0]
		}
		return -1
	}
	return list.heights[index]
}

//AddNewPair adding new pair to list
func (list *IotexList) AddNewPair(height int, time int) bool {
	//fmt.Printf("Add %d and %d\n", height, time)
	index := list.validate(height, time)

	if index < -1 {
		return false
	}

	list.addToSlie(index, height, time)
	list.addToHashTable(height, time)
	return true
}

func (list *IotexList) addToSlie(index int, height int, time int) {
	//fmt.Printf("Add %d and %d index=%d size=%d\n", height, time, index, len(list.heights))
	if index != len(list.heights) {
		index = index + 1
	}
	//fmt.Printf("Add %d and %d index=%d size\n", height, time, index)
	if len(list.heights) == 0 {
		list.heights = append(list.heights, height)
		list.times = append(list.times, time)
	} else {
		list.heights = append(list.heights, 0)             // Step 1
		copy(list.heights[index+1:], list.heights[index:]) // Step 2
		list.heights[index] = height
		list.times = append(list.times, 0)             // Step 1
		copy(list.times[index+1:], list.times[index:]) // Step 2
		list.times[index] = time

		// tmp := append(list.heights[:index], height)
		// list.heights = append(tmp, list.heights[index+1:]...)
		// tmp = append(list.times[:index], time)
		// list.times = append(tmp, list.times[index+1:]...)
	}
}

func (list *IotexList) addToHashTable(height int, time int) {
	list.tableHeighs[height] = true
	list.tableTimes[time] = true
}

func (list *IotexList) validate(height int, time int) int {
	if list.isConflict() {
		return -4
	}
	if len(list.heights) == 0 && len(list.times) == 0 {
		//fmt.Printf("here list =0\n")
		return 0
	}
	if list.checkContainHeight(height) || list.checkContainTime(time) {
		return -3
	}
	heightIndex := list.binarySearch(list.heights, height, 0, len(list.heights)-1)
	timeIndex := list.binarySearch(list.times, time, 0, len(list.times)-1)

	//fmt.Printf("height index=%d\ttime index=%d\n", heightIndex, timeIndex)
	if heightIndex != timeIndex {
		return -2
	}
	return heightIndex
}

func (list *IotexList) binarySearch(data []int, value int, start int, end int) int {
	mid := (start + end) / 2

	// Need to go lower but can't
	if mid == start && data[mid] > value {
		return -1
	}

	// Found a candidate, and can't go higher
	if data[mid] <= value && (mid == end || (mid != len(data)-1 && data[mid+1] > value)) {
		//System.out.println("mid="+mid+"\tvalue="+value+"\tdata="+data.get(mid));
		//fmt.Printf("mid=%d value=%d data=%d\n", mid, value, data[mid])
		return mid
	}

	if data[mid] <= value {
		// Consider upper half
		return list.binarySearch(data, value, mid+1, end)
	} else {
		// Consider lower half
		return list.binarySearch(data, value, start, mid)
	}
}

func (list *IotexList) isConflict() bool {
	if len(list.heights) != len(list.times) {
		return true
	}
	return false
}

func Init() *IotexList {
	list := new(IotexList)
	list.heights = make([]int, 0)
	list.times = make([]int, 0)
	list.tableHeighs = make(map[int]bool)
	list.tableTimes = make(map[int]bool)
	return list
}

/*
func main() {
	list := Init()
	list.AddNewPair(1, 1)
	list.printPair()
	list.AddNewPair(2, 2)
	list.printPair()
	list.AddNewPair(4, 2)
	list.printPair()
	list.AddNewPair(5, 5)
	list.printPair()
	list.AddNewPair(1, 6)
	list.printPair()
	list.AddNewPair(4, 4)
	list.printPair()
	list.AddNewPair(3, 3)
	list.printPair()
	list.AddNewPair(0, 0)
	list.printPair()
	list.AddNewPair(6, 7)
	list.printPair()
	list.AddNewPair(7, 6)
	list.printPair()
	list.AddNewPair(7, 8)
	list.printPair()
	fmt.Printf("%v", list.GetList())

	fmt.Printf("query with time: -1\theight: %d\n", list.nearestHeightBefore(-1))
	fmt.Printf("query with time: 0\theight: %d\n", list.nearestHeightBefore(0))
	fmt.Printf("query with time: 1\theight: %d\n", list.nearestHeightBefore(1))
	fmt.Printf("query with time: 2\theight: %d\n", list.nearestHeightBefore(2))
	fmt.Printf("query with time: 3\theight: %d\n", list.nearestHeightBefore(3))
	fmt.Printf("query with time: 4\theight: %d\n", list.nearestHeightBefore(4))
	fmt.Printf("query with time: 5\theight: %d\n", list.nearestHeightBefore(5))
	fmt.Printf("query with time: 6\theight: %d\n", list.nearestHeightBefore(6))
	fmt.Printf("query with time: 7\theight: %d\n", list.nearestHeightBefore(7))
	fmt.Printf("query with time: 8\theight: %d\n", list.nearestHeightBefore(8))
	fmt.Printf("query with time: 9\theight: %d\n", list.nearestHeightBefore(9))
	fmt.Printf("query with time: 10\theight: %d\n", list.nearestHeightBefore(10))

}
*/
func compare(list1 []Pair, list2 []Pair) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i, pair := range list1 {
		if pair.height != list2[i].height || pair.time != list2[i].time {
			return false
		}
	}

	return true
}
