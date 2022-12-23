package plus_one

func addBinary(a string, b string) string {

	i := len(a) - 1
	j := len(b) - 1

	b0 := '0'
	b1 := '1'

	var va, vb, mv rune
	mv = b0
	var s string

	for i >= 0 || j >= 0 {

		if i >= 0 {
			va = rune(a[i])
			i--
		} else {
			va = b0
		}

		if j >= 0 {
			vb = rune(b[j])
			j--
		} else {
			vb = b0
		}

		if (va == b0 && vb == b0) && mv == b1 {
			s = "1" + s
			mv = b0
		} else if (va == b1 && vb == b1) && mv == b1 {
			s = "1" + s
			mv = b1
		} else if (va == b1 && vb == b1) && mv == b0 {
			s = "0" + s
			mv = b1
		} else if (va == b1 || vb == b1) && mv == b1 {
			s = "0" + s
			mv = b1
		} else if (va == b1 || vb == b1) && mv == b0 {
			s = "1" + s
			mv = b0
		} else {
			mv = b0
			s = "0" + s
		}

	}
	if mv == b1 {
		return "1" + s
	}
	return s
}
