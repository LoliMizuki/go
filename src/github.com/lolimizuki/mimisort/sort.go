//  The way to Go 11.7 範例
package mimisort

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i+1, i)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			return false
		}
	}

	return true
}
