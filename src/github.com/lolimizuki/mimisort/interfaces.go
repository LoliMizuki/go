//  The way to Go 11.7 範例
package mimisort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
