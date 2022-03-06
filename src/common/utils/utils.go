package utils

import (
	"fmt"
	"net"
	"time"
)

func TestTCPConn(addr string, timeout, interval int) error {
	success := make(chan int, 1)
	cancel := make(chan int, 1)

	go func() {
		n := 1

	loop:
		for {
			select {
			case <-cancel:
				break
			default:
				conn, err := net.DialTimeout("tcp", addr, time.Duration(n)*time.Second)
				if err != nil {
					n = n * 2
					time.Sleep(time.Duration(interval) * time.Second)
					continue
				}
				if err = conn.Close(); err != nil {
					fmt.Print("hi")
				}
				success <- 1
				break loop
			}
		}
	}()

	select {
	case <-success:
		return nil
	case <-time.After(time.Duration(timeout) * time.Second):
		cancel <- 1
		return fmt.Errorf("failed to connect to tcp:%s after %d seconds", addr, timeout)
	}
}
