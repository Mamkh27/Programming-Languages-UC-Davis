package bug1
import(
	"sync"
)
var m sync.Mutex
// Counter stores a count.
type Counter struct {
	n int64
}

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	m.Lock()
	c.n++
	m.Unlock()
}