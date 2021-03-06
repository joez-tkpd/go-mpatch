// +build 386

package mpatch

// Gets the jump function rewrite bytes
func getJumpFuncBytes(to uintptr) ([]byte, error) {
	return []byte{
		0xBA,
		byte(to),
		byte(to >> 8),
		byte(to >> 16),
		byte(to >> 24),
		0xFF, 0x22,
	}, nil
}
