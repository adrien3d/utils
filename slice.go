package utils

func RemoveStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func contains(elts []interface{}, elt interface{}) bool {
	for _, element := range elts {
		if element == elt {
			return true
		}
	}
	return false
}

func containsInt(elts []int, elt int) bool {
	s := make([]interface{}, len(elts))
	for i, v := range elts {
		s[i] = v
	}
	return contains(s, elt)
}

func containsFloat(elts []int, elt float64) bool {
	s := make([]interface{}, len(elts))
	for i, v := range elts {
		s[i] = v
	}
	return contains(s, elt)
}

func containsString(elts []int, elt string) bool {
	s := make([]interface{}, len(elts))
	for i, v := range elts {
		s[i] = v
	}
	return contains(s, elt)
}

func containsBool(elts []int, elt bool) bool {
	s := make([]interface{}, len(elts))
	for i, v := range elts {
		s[i] = v
	}
	return contains(s, elt)
}
