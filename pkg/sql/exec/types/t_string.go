// Code generated by "stringer -type=T"; DO NOT EDIT.

package types

import "strconv"

const _T_name = "BoolBytesInt8Int16Int32Int64Float32Float64Unhandled"

var _T_index = [...]uint8{0, 4, 9, 13, 18, 23, 28, 35, 42, 51}

func (i T) String() string {
	if i < 0 || i >= T(len(_T_index)-1) {
		return "T(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _T_name[_T_index[i]:_T_index[i+1]]
}