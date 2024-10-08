package ftpserver

import (
	"errors"
	"time"
)

type portAllocator struct {
	ports chan int
}

func newPassivePortAllocator(portRange *PortRange) *portAllocator {
	ports := make(chan int, portRange.End-portRange.Start)
	for i := portRange.Start; i <= portRange.End; i++ {
		ports <- i
	}
	return &portAllocator{ports}
}

func (a *portAllocator) GetPort() (int, error) {
	select {
	case port := <-a.ports:
		return port, nil
	case <-time.After(time.Second * 10):
		break
	}
	return 0, errors.New("timeout")
}

func (a *portAllocator) ReleasePort(port int) {
	a.ports <- port
}
