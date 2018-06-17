package diff

import (
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	testCases := []struct {
		list1 []string
		list2 []string
		want  []string
	}{
		{
			list1: []string{"aa"},
			list2: []string{},
			want:  []string{"-aa"},
		},
		{
			list1: []string{"aa", "bb"},
			list2: []string{},
			want:  []string{"-aa", "-bb"},
		},
		{
			list1: []string{},
			list2: []string{"aa"},
			want:  []string{"+aa"},
		},
		{
			list1: []string{},
			list2: []string{"aa", "bb"},
			want:  []string{"+aa", "+bb"},
		},
		{
			list1: []string{"aa", "bb", "cc", "dd", "ee"},
			list2: []string{"aa", "aa", "bb", "dd", "ff", "gg"},
			want:  []string{"+aa", "-cc", "+ff", "+gg", "-ee"},
		},
		{
			list1: []string{"aa", "bb", "cc", "dd", "ee", "hh"},
			list2: []string{"aa", "aa", "bb", "dd", "ff", "gg"},
			want:  []string{"+aa", "-cc", "+ff", "+gg", "-ee", "-hh"},
		},
		{
			list1: []string{"aa", "bb", "cc", "dd", "ll", "ee", "hh"},
			list2: []string{"aa", "aa", "bb", "dd", "ll", "ff", "gg", "cc"},
			want:  []string{"+aa", "-cc", "+ff", "+gg", "+cc", "-ee", "-hh"},
		},
		{
			list1: []string{"aa", "bb", "cc", "ll", "dd", "ee", "hh"},
			list2: []string{"aa", "aa", "bb", "dd", "ff", "gg", "cc", "ll"},
			want:  []string{"+aa", "+dd", "+ff", "+gg", "-dd", "-ee", "-hh"},
		},
	}

	for _, tc := range testCases {
		got := Compare(tc.list1, tc.list2)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("incorrect diff between %v and %v, got: %v, want: %v", tc.list1, tc.list2, got, tc.want)
		}
	}
}

func TestCompareFile(t *testing.T) {
	testCases := []struct {
		file1 string
		file2 string
		want  []string
	}{
		{
			file1: "testfile1.txt",
			file2: "testfile2.txt",
			want: []string{
				"-echo", "+hello", "+bus", "-hotel", "-november",
				"-siera", "+romeo", "+hotel", "+potato", "-victor",
				"+computer", "+alexa"},
		},
	}

	for _, tc := range testCases {
		got, err := CompareFiles(tc.file1, tc.file2)
		if err != nil {
			t.Errorf("error comparing files: %v", err)
			continue
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("incorrect diff between %v and %v, got: %v, want: %v", tc.file1, tc.file2, got, tc.want)
		}
	}
}
