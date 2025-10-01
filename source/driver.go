package main

import (
	// "errors"
	// "fmt"

	// "strconv"
	// "strings"
	// "time"

	"os/exec"
	"strings"

	//"sync"
	"time"

	"github.com/Dartmouth-OpenAV/microservice-framework/framework"
)

func pingHost(socketKey string) (string, error) {
	// function := "pingHost"

	value := `"unknown"`
	err := error(nil)
	maxRetries := 2
	for maxRetries > 0 {
		value, err = pingHostDo(socketKey)
		if value == `"unknown"` { // Something went wrong - perhaps try again
			maxRetries--
			time.Sleep(1 * time.Second)
		} else { // Succeeded
			maxRetries = 0
		}
	}

	return value, err
}

func pingHostDo(socketKey string) (string, error) {
	function := "pingHostDo"
	host := ""
	var cmd *exec.Cmd

	// Strip the username, password, and port number from the socket key
	// Find the first occurrence of ":"
	firstAtSymbol := strings.Index(socketKey, "@")
	if firstAtSymbol == -1 {
		framework.AddToErrors(socketKey, function+" - 2q34awev6 no at symbol found in sockeyKey"+socketKey) // No colon found
	}

	// Find the last occurrence of ":"
	lastColon := strings.LastIndex(socketKey, ":")
	if lastColon <= firstAtSymbol {
		framework.AddToErrors(socketKey, function+" - 2q34awev7 only one colon in socketKey"+socketKey) // Only one colon or last colon is before first
	}
	
	value := `"unknown"`
	// Extract the substring between firstAtSymbol+1 and lastColon
	if firstAtSymbol != -1 && lastColon > firstAtSymbol {
		host = socketKey[firstAtSymbol+1 : lastColon]
		cmd = exec.Command("ping", "-c", "1", "-W", "1", host)
		err := cmd.Run()		
		if err != nil {
			framework.AddToErrors(socketKey, function+" - 2q34awev5 ping "+host+" failed")
			errMsg := 
			value = `"false"`
		} else { // Succeeded
			value = `"true"`
		}
	}
	// If we got here, the response was good, so successful return with the state indication
	return value, nil
}
