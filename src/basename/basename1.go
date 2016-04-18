package basename1

// basename ...

// a => a, a.go => a.go. a/b/c.go => c.go
func basename(s string) string  {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] = '/' {
			s = s[i+1:]
			break
		}
	}
	// Oreserve everything before last '.'
	for i := len(s) -1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}