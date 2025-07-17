package sequencecompare

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func GenerateSequence(start, finish int) string {
	var items []int

	for i := start; i < finish+1; i++ {
		items = append(items, i)
	}

	return PrintSlice(items)
}

func PrintSlice(items []int) string {
	var strItems []string

	for _, item := range items {
		strItems = append(strItems, strconv.Itoa(item))
	}

	return fmt.Sprintf("var x = []int{%s}", strings.Join(strItems, ", "))
}

func RandomSortItems(items []int) []int {
	for i := len(items) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		items[i], items[j] = items[j], items[i]
	}

	return items
}
