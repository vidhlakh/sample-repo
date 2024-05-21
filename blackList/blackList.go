package blackList

import (
	"context"
	"fmt"
	"time"

	"github.com/vidhlakh/sample-repo/models"
)

func Blacklist(ctx context.Context, input models.Input, blCh chan bool) {
	subBlCh := make(chan bool)
	// if any of the subrule return true, then return true
	go subRule1(input, subBlCh)
	//go subRule2(input, subBlCh)
	for blRes := range subBlCh {
		if !blRes {
			fmt.Println("Send from blacklist rule", blRes)
			blCh <- blRes
			ctx.Done()
		}
	}

}

func subRule1(input models.Input, subBlCh chan<- bool) {
	time.Sleep(30 * time.Second)
	if input.City == "XXX" {
		subBlCh <- true // blacklist the input if city is XXX
	} else {
		subBlCh <- false
	}
}

func subRule2(input models.Input, subBlCh chan<- bool) {
	time.Sleep(1 * time.Minute)

	if input.Country == "XXX" {
		subBlCh <- true
	} else {
		subBlCh <- false
	}
}
