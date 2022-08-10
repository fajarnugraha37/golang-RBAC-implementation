package util

func Map[T, U any](inputs []T, callback func(item T) U) []U {
	var output []U
	for _, item := range inputs {
		output = append(output, callback(item))
	}
	return output
}

func InArray[T ~byte | ~string | ~int64 | ~float64 | ~bool](target T, items []T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func Intersection[T ~byte | ~string | ~int64 | ~float64 | ~bool](s1, s2 []T) (inter []T) {
	hash := make(map[T]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	inter = RemoveDups(inter)
	return
}

func RemoveDups[T ~byte | ~string | ~int64 | ~float64 | ~bool](elements []T) (nodups []T) {
	encountered := make(map[T]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

func RemoveEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
