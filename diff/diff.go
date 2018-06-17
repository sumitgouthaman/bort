package diff

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func CompareFiles(file1, file2 string) ([]string, error) {
	f1Bytes, err := ioutil.ReadFile(file1)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file %v: %v", file1, err)
	}
	f2Bytes, err := ioutil.ReadFile(file2)
	if err != nil {
		return nil, fmt.Errorf("couldn't open file %v: %v", file2, err)
	}

	f1Contents := string(f1Bytes)
	f2Contents := string(f2Bytes)

	f1Lines := strings.Split(f1Contents, "\n")
	f2Lines := strings.Split(f2Contents, "\n")

	return Compare(f1Lines, f2Lines), nil
}

func Compare(list1 []string, list2 []string) []string {
	return compare(list1, list2, make(map[pos][]string), pos{list1Pos: 0, list2Pos: 0})
}

type pos struct {
	list1Pos int
	list2Pos int
}

func compare(list1 []string, list2 []string, cache map[pos][]string, p pos) []string {
	d, ok := cache[p]
	if ok {
		return duplicate(d)
	}

	if len(list1) == 0 {
		out := applyPrefix(list2, "+")
		cache[p] = out
		return duplicate(out)
	}
	if len(list2) == 0 {
		out := applyPrefix(list1, "-")
		cache[p] = out
		return duplicate(out)
	}

	if list1[0] == list2[0] {
		out := compare(list1[1:], list2[1:], cache, pos{list1Pos: p.list1Pos + 1, list2Pos: p.list2Pos + 1})
		cache[p] = out
		return duplicate(out)
	}

	caseAdded := compare(list1, list2[1:], cache, pos{list1Pos: p.list1Pos, list2Pos: p.list2Pos + 1})
	caseDeleted := compare(list1[1:], list2, cache, pos{list1Pos: p.list1Pos + 1, list2Pos: p.list2Pos})

	if len(caseAdded) > len(caseDeleted) {
		out := append([]string{"-" + list1[0]}, caseDeleted...)
		cache[p] = out
		return duplicate(out)
	}

	out := append([]string{"+" + list2[0]}, caseAdded...)
	cache[p] = out
	return duplicate(out)
}

func duplicate(arr []string) []string {
	tmp := make([]string, len(arr))
	copy(tmp, arr)
	return tmp
}

func applyPrefix(src []string, prefix string) []string {
	out := make([]string, 0)
	for _, s := range src {
		out = append(out, prefix+s)
	}
	return out
}
