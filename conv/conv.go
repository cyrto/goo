package conv

import "strconv"

//interface to string
func Interface2String(inter interface{}) (res string) {
	switch inter.(type) {
	case string:
		res = inter.(string)
		break
	}
	return
}

//string to int64
func String2Int64(s string) (res int64) {
	i, _ := strconv.Atoi(s)
	return int64(i)
}

//interface to int64
func Interface2Int64(inter interface{}) int64 {
	var res int64
	switch inter.(type) {
	case int64:
		res = inter.(int64)
		break
	}
	return res
}

//interface to float64
func Interface2Float64(inter interface{}) float64 {
	var res float64
	switch inter.(type) {
	case float64:
		res = inter.(float64)
		break
	}
	return res
}
