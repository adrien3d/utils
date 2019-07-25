package utils

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