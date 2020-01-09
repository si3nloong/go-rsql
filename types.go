package rsql

// Strings :
type Strings []string

// IndexOf :
func (slice Strings) IndexOf(search string) (idx int) {
	idx = -1
	length := len(slice)
	for i := 0; i < length; i++ {
		if slice[i] == search {
			idx = i
			break
		}
	}
	return
}
