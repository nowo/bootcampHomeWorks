package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxConcurrentUser = 2

var wg sync.WaitGroup // 1

func main() {
	wg.Add(1)
	waitChan := make(chan struct{}, maxConcurrentUser)
	l := &License{}

	// Start a goroutine for each user
	waitChan <- struct{}{}
	for {
		wg.Add(1)
		go func() {
			// Wait for a user to take a license
			err := takeLicense(l, maxConcurrentUser)
			if err != nil {
				fmt.Println(err)
			}
			// Signal a user to take a license
			<-waitChan
			// Wait for a user to finish
			wg.Done()
		}()
		waitChan <- struct{}{}
	}
	wg.Wait()
}

// NoLicenseError is a custom error type
type NoLicenseError struct{}

func (m *NoLicenseError) Error() string {
	return "Maximum number of concurrent user reached"
}

// License is a struct that contains the number of concurrent user
type License struct {
	numberOfTakenLicense int
}

// takeLicense is a function that takes a license
func (l *License) takeLicense() int {
	l.numberOfTakenLicense = l.numberOfTakenLicense + 1
	return l.numberOfTakenLicense
}

// takeLicense is a function that takes a license and returns an error if the number of concurrent user is reached
func takeLicense(l *License, maxConcurrentUser int) error {
	defer wg.Done()

	if l.numberOfTakenLicense >= maxConcurrentUser {
		return &NoLicenseError{}
	}
	l.takeLicense()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Println("User took license")
	reduceLicense(&l.numberOfTakenLicense)
	return nil

}

func reduceLicense(numberOfTakenLicense *int) {
	*numberOfTakenLicense = *numberOfTakenLicense - 1
}
