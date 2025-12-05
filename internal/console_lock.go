package internal

import "sync"

// ConsoleLock serializes writes to the terminal to avoid interleaved output
var ConsoleLock sync.Mutex
