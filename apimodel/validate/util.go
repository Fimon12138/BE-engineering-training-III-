package validate

func CheckPages(pageNo, pageSize int) bool {
	if pageNo <= 0 || pageSize <= 0 {
		return false
	}

	return true
}
