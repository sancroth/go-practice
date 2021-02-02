package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	var emptySlice []string
	s.Show("emptySlice", emptySlice)

	emptySlice = []string{}
	// all empty slices in go point to the same dummy slice(ptr 168)
	s.Show("emptySlice", emptySlice)

	langSlice := []string{"javascript", "java", "go", "ruby", "python"}

	s.Show("langSlice", langSlice)

	partSlice := langSlice[:0]
	// 0 length cap(langSlice) capacity
	// accessing it directly returns no available elements
	s.Show("partSlice", partSlice)
	s.Show("partSlice[:0]", partSlice[:0])
	// reslicing gives access back to the elements
	// since they become available again(they always existed)
	s.Show("partSlice[:cap]", partSlice[:cap(partSlice)])

	// simple copy just for the sake of it
	backupSlice := make([]string, cap(partSlice))
	copy(backupSlice, langSlice)

	s.Show("backupSlice", backupSlice)

	// reducing the slice capacity will lead on the
	// ptr to move up by unsafe.Sizeof("") bytes <- string
	for cap(partSlice) > 0 {
		partSlice = partSlice[1:cap(partSlice)]
		s.Show("partSlice", partSlice)
	}

	s.Show("backupSlice", backupSlice)

}
