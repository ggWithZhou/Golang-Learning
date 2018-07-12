package time

import (
	"fmt"
	"testing"
	"time"
)

func TestStringVar(t *testing.T) {
	const layout = "Portscan-5"
	ti := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	fmt.Println(ti.Format(layout))
	fmt.Println(ti.UTC().Format(layout))
}
