// Code generated by "stringer -type=KaleidoToken"; DO NOT EDIT.

package lexer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KTokenEOF-0]
	_ = x[KTokenDef-1]
	_ = x[KTokenExtern-2]
	_ = x[KTokenIdentifier-3]
	_ = x[KTokenNumber-4]
	_ = x[KTokenSymbol-5]
}

const _KaleidoToken_name = "KTokenEOFKTokenDefKTokenExternKTokenIdentifierKTokenNumberKTokenSymbol"

var _KaleidoToken_index = [...]uint8{0, 9, 18, 30, 46, 58, 70}

func (i KaleidoToken) String() string {
	if i < 0 || i >= KaleidoToken(len(_KaleidoToken_index)-1) {
		return "KaleidoToken(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _KaleidoToken_name[_KaleidoToken_index[i]:_KaleidoToken_index[i+1]]
}
