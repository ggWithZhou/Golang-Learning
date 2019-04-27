import testing
import time

func TestForTest(t *testing.T){
  go fmt.Println("test")
  time.Sleep(time.Second)
}
