package whiteList

import (
	"context"
	"fmt"

	"github.com/vidhlakh/sample-repo/models"
)

func Whitelist(ctx context.Context, input models.Input, whCh chan bool) {
	//result := false
	subWhCh := make(chan bool)
	//ctx, cancel := context.WithCancel(ctx)
	// if any of the subrule return true, then return true

	go subRule1(input, subWhCh)
	//go subRule2(input, subWhCh)
	for whRes := range subWhCh {
		if whRes {
			fmt.Println("Send from whitelist rule", whRes)
			whCh <- whRes
			ctx.Done()
		}
	}
	//cancel()
}

func subRule2(input models.Input, subWhCh chan<- bool) {
	//time.Sleep(1 * time.Minute)
	if input.City == "Chennai" || input.City == "Bangalore" {
		subWhCh <- true
	} else {
		subWhCh <- false
	}
}

func subRule1(input models.Input, subWhCh chan<- bool) {

	if input.Country == "India" {
		subWhCh <- true
	} else {
		subWhCh <- false
	}
}
