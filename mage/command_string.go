// Code generated by "stringer -type=Command"; DO NOT EDIT.

package mage

import "fmt"

const _Command_name = "NONEVERSIONINITCLEAN"

var _Command_index = [...]uint8{0, 4, 11, 15, 20}

func (i Command) String() string {
	if i < 0 || i >= Command(len(_Command_index)-1) {
		return fmt.Sprintf("Command(%d)", i)
	}
	return _Command_name[_Command_index[i]:_Command_index[i+1]]
}
