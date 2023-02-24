package main

import "fmt"

// 背包问题

type Item struct {
	weight int
	price  int
	name   string
}

func main() {
	items := []Item{
		{
			weight: 0,
			price:  0,
			name:   "empty",
		},
		{
			weight: 1,
			price:  6,
			name:   "C",
		},
		{
			weight: 3,
			price:  10,
			name:   "A",
		},
		{
			weight: 3,
			price:  13,
			name:   "D",
		},
		{
			weight: 2,
			price:  9,
			name:   "B",
		},

		{
			weight: 3,
			price:  11,
			name:   "F",
		},
	}
	bagV := 6

	record := make([][][]Item, len(items))
	for i := 0; i < len(items); i++ {
		record[i] = make([][]Item, bagV+1)
	}

	for bagWeight := 0; bagWeight < bagV+1; bagWeight++ {
		record[0][bagWeight] = []Item{
			{
				weight: 0,
				price:  0,
				name:   "empty",
			},
		}
	}

	for itemIndex := 0; itemIndex < len(items); itemIndex++ {
		record[itemIndex][0] = []Item{
			{
				weight: 0,
				price:  0,
				name:   "empty",
			},
		}
	}

	for bagWeight := 1; bagWeight < bagV+1; bagWeight++ {
		for woodIndex := 1; woodIndex < len(items); woodIndex++ {
			if items[woodIndex].weight > bagWeight {
				src := record[woodIndex-1][bagWeight]
				dst := make([]Item, len(src))
				copy(dst, src)
				record[woodIndex][bagWeight] = dst
				//record[itemIndex][bagWeight] = record[itemIndex-1][bagWeight]
			} else {
				ws1 := make([]Item, len(record[woodIndex-1][bagWeight]))
				copy(ws1, record[woodIndex-1][bagWeight])
				//ws1 := record[itemIndex-1][bagWeight]

				temp := make([]Item, len(record[woodIndex-1][bagWeight-items[woodIndex].weight]))
				copy(temp, record[woodIndex-1][bagWeight-items[woodIndex].weight])
				ws2 := append(temp, items[woodIndex])
				//ws2 := append(record[itemIndex-1][bagWeight-items[itemIndex].weight], items[itemIndex])
				if Sum(ws1) > Sum(ws2) {
					record[woodIndex][bagWeight] = ws1
				} else {
					record[woodIndex][bagWeight] = ws2
				}
			}
		}
	}

	fmt.Println(record[len(items)-1][bagV])
}

func Sum(s []Item) (sum int) {
	for _, v := range s {
		sum += v.price
	}
	return
}
