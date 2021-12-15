package main

import (
	
	"errors"
	"fmt"
	/*
	"log"
	*/
	"time"
	"github.com/coinbase/rosetta-sdk-go/fetcher"
	/*
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	*/
)

func main() {
	
	DefaultURL                               := "http://localhost:8080"
	DefaultTimeout                           := 10
	DefaultMaxRetries                        := 5
	DefaultMaxOnlineConnections              := 120 // most OS have a default limit of 128
	/*
	DefaultMaxOfflineConnections             := 4   // we shouldn't need many connections for construction
	DefaultMaxSyncConcurrency                := 64
	DefaultActiveReconciliationConcurrency   := 16
	DefaultInactiveReconciliationConcurrency := 4
	DefaultInactiveReconciliationFrequency   := 250
	DefaultConfirmationDepth                 := 10
	DefaultStaleDepth                        := 30
	DefaultBroadcastLimit                    := 3
	DefaultTipDelay                          := 300
	DefaultBlockBroadcastLimit               := 5
	DefaultStatusPort                        := 9090
	DefaultMaxReorgDepth                     := 100
	*/

	MaxOnlineConnections					 := DefaultMaxOnlineConnections
	RetryElapsedTime                         := 5
	HTTPTimeout 							 := DefaultTimeout
	var MaxRetries uint64                    =  uint64(DefaultMaxRetries)
	OnlineURL								 := DefaultURL
	fetcherOpts := []fetcher.Option{
		fetcher.WithMaxConnections(MaxOnlineConnections),
		fetcher.WithRetryElapsedTime(time.Duration(RetryElapsedTime) * time.Second),
		fetcher.WithTimeout(time.Duration(HTTPTimeout) * time.Second),
		fetcher.WithMaxRetries(MaxRetries),
	}


	f := fetcher.New(
		OnlineURL,
		fetcherOpts...,
	)

	networkList, fetchErr := f.NetworkListRetry(Context, nil)
	if fetchErr != nil {
		fmt.Errorf("%w: unable to fetch network list", fetchErr.Err)
		return 
	}

	if len(networkList.NetworkIdentifiers) == 0 {
		errors.New("no networks available")
		return
	}
}