package Carbon

import "time"

// Sleep sleep()
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}
func USleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
