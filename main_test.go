package main

import (
	"testing"

	"github.com/vidhlakh/sample-repo/models"
)

func TestRuleEngineWhitelist(t *testing.T) {
	type args struct {
		input models.Input
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "rules from Whitelist is passing",
			args: args{
				input: models.Input{
					Name:       "cust1",
					CardNumber: "1234",
					City:       "XXX",
					Address:    "1st cross street , Guindy",
					Country:    "India",
					Merchant:   "Corporate",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RuleEngine(tt.args.input)
		})
	}
}

func TestRuleEngineBlacklist(t *testing.T) {
	type args struct {
		input models.Input
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "rules from Blacklist is passing",
			args: args{
				input: models.Input{
					Name:     "cust1",
					City:     "Chennai",
					Address:  "1st cross street , Guindy",
					Country:  "Ind",
					Merchant: "Corporate",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RuleEngine(tt.args.input)
		})
	}
}

func TestRuleEngineMerchant(t *testing.T) {
	type args struct {
		input models.Input
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "rules from Merchant is passing",
			args: args{
				input: models.Input{
					Name:     "cust1",
					City:     "XXX",
					Address:  "1st cross street , Guindy",
					Country:  "Dubai",
					Merchant: "Individual",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RuleEngine(tt.args.input)
		})
	}
}

func TestRuleEngineDefault(t *testing.T) {
	type args struct {
		input models.Input
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "rules from Default is passing",
			args: args{
				input: models.Input{
					Name:       "cust1",
					CardNumber: "1234",
					City:       "XXX",
					Address:    "1st cross street , Guindy",
					Country:    "USA",
					Merchant:   "Corporate",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RuleEngine(tt.args.input)
		})
	}
}
