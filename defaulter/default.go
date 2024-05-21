package defaulter

import (
	"context"
	"fmt"

	"github.com/vidhlakh/sample-repo/models"
)

func Default(ctx context.Context, input models.Input, defCh chan bool) {
	//result := false
	subdefCh := make(chan bool)
	//ctx, cancel := context.WithCancel(ctx)
	// if any of the subrule return true, then return true

	go subRule1(input, subdefCh)
	//go subRule2(input, subdefCh)
	for defRes := range subdefCh {
		if defRes {
			fmt.Println("Send from default rule", defRes)
			defCh <- defRes
			ctx.Done()
		}
	}
	//cancel()
}

func subRule1(input models.Input, subdefCh chan<- bool) {
	//time.Sleep(10 * time.Second)
	if input.CardNumber != "" {
		subdefCh <- true
	} else {
		subdefCh <- false
	}
}
