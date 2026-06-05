// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import "sync"

// PanicRecoveredMonitorFunc reports a panic recovered by Gin.
type PanicRecoveredMonitorFunc func(component, operation string)

var panicRecoveredMonitor = struct {
	sync.RWMutex
	fn PanicRecoveredMonitorFunc
}{}

// SetPanicRecoveredMonitor sets the process-level reporter used by Recovery middleware.
func SetPanicRecoveredMonitor(fn PanicRecoveredMonitorFunc) {
	panicRecoveredMonitor.Lock()
	defer panicRecoveredMonitor.Unlock()

	panicRecoveredMonitor.fn = fn
}

func reportPanicRecovered(component, operation string) {
	panicRecoveredMonitor.RLock()
	fn := panicRecoveredMonitor.fn
	panicRecoveredMonitor.RUnlock()

	if fn == nil {
		return
	}

	defer func() {
		_ = recover()
	}()
	fn(component, operation)
}
