package utils

func CheckNullString(str string) *string {
	if len(str) == 0 {
		return nil
	}

	return &str
}

func RefStringArr(s []string) []*string {
	if s != nil {
		ptrs:=make([]*string,len(s))
		for  i := 0; i < len(s); i++ {
			ptrs[i] = &s[i] /* assign the address of string. */
		}
		return ptrs
	}
	return nil
}