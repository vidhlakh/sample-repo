package main

import (
	"context"
	"fmt"

	"github.com/vidhlakh/sample-repo/blackList"
	"github.com/vidhlakh/sample-repo/defaulter"

	"github.com/vidhlakh/sample-repo/merchant"
	"github.com/vidhlakh/sample-repo/models"
	"github.com/vidhlakh/sample-repo/whiteList"
)

func main() {
	fmt.Println("interview")
	input := models.Input{
		Name:       "cust1",
		CardNumber: "1234",
		City:       "Chennai",
		Address:    "1st cross street , Guindy",
		Country:    "India",
		Merchant:   "Corporate",
	}
	RuleEngine(input)
}

func RuleEngine(input models.Input) {
	whCh := make(chan bool)
	blCh := make(chan bool)
	merCh := make(chan bool)
	defCh := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go whiteList.Whitelist(ctx, input, whCh)
	go blackList.Blacklist(ctx, input, blCh)
	go merchant.Merchant(ctx, input, merCh)
	go defaulter.Default(ctx, input, defCh)

	for {
		select {
		case wh, ok := <-whCh:
			if !ok {
				ctx.Done()
				return
			}
			if wh {
				fmt.Println("Result from whitelist is ", wh)
				cancel()
			}
		case bl, ok := <-blCh:
			if !ok {
				ctx.Done()
				return
			}
			if !bl {
				fmt.Println("Result from blackist is ", bl)
				cancel()
			}
		case mer, ok := <-merCh:
			if !ok {
				ctx.Done()
				return
			}
			if mer {
				fmt.Println("Result from merchant is ", mer)
				cancel()
			}
		case def, ok := <-defCh:
			if !ok {
				ctx.Done()
				return
			}
			if def {
				fmt.Println("Result from default is ", def)
				cancel()
			}
		case <-ctx.Done():
			fmt.Println("Completed")
			return
		}
	}
}

/*
Rule Engine - input can be anything (name, city, cardnumber, Address) - 50 input params
Output - true / false
Rule engine
- 4 groups
Mandate whitelist
    - if any rules pass - return true. If any subgroup rule pass , then we can come out
    - sub rules (static and dynamic rules [db/api calls] )
Mandate Blacklist
   - if any rules pass - return false
Merchant
   - if any rules pass - return true
Default
    - if any rules pass - return true
Follow hierarchy in order of groups

1. Http server
2. API
3. Service with DB , external API inte
4. Folder File reading logic
5. Goroutine for the group call the goroutin in subgroups
and exist using signal
channel select ()


*/
