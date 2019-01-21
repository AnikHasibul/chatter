package util

// SafeClose closes a channel so rudely but safely.
func SafeClose(ch chan bool) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()
	close(ch)
	return true
}

// SafeClose sendss a value to a channel so rudely but safely.
func SafeSend(ch chan bool, value bool) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = true
		}
	}()
	ch <- value
	return false
}
