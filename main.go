package main

import (
	/*
	"errors"
	"fmt"
	"log"
	*/
	"time"
	"github.com/coinbase/rosetta-sdk-go/fetcher"
	"github.com/keninqiu/rosetta-test/configurations/types"
	/*
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	*/
)

func main() {
	
	fetcherOpts := []fetcher.Option{
		fetcher.WithMaxConnections(Config.MaxOnlineConnections),
		fetcher.WithRetryElapsedTime(time.Duration(Config.RetryElapsedTime) * time.Second),
		fetcher.WithTimeout(time.Duration(Config.HTTPTimeout) * time.Second),
		fetcher.WithMaxRetries(Config.MaxRetries),
	}
/*
	if Config.ForceRetry {
		fetcherOpts = append(fetcherOpts, fetcher.WithForceRetry())
	}

	f := fetcher.New(
		Config.OnlineURL,
		fetcherOpts...,
	)
	*/

}