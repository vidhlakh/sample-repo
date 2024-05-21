package merchant

import (
	"context"
	"fmt"

	"github.com/vidhlakh/sample-repo/models"
)

func Merchant(ctx context.Context, input models.Input, merCh chan bool) {
	//result := false
	subMerCh := make(chan bool)
	//ctx, cancel := context.WithCancel(ctx)
	// if any of the subrule return true, then return true

	go subRule1(input, subMerCh)
	//go subRule2(input, subMerCh)
	for merRes := range subMerCh {
		if merRes {
			fmt.Println("Send from merchant rule", merRes)
			merCh <- merRes
			ctx.Done()
		}
	}
	//cancel()
}

func subRule1(input models.Input, subMerCh chan<- bool) {
	//time.Sleep(20 * time.Second)
	if input.Merchant == "Individual" {
		subMerCh <- true
	} else {
		subMerCh <- false
	}
}
