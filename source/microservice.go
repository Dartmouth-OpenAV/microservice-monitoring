package main

import (
	"errors"

	"github.com/Dartmouth-OpenAV/microservice-framework/framework"
)

// Define device specific globals
// var defaultLineDelimiter = 0xA // 0xA==\n==linefeed, 0xD==\r==carriagereturn

func setFrameworkGlobals() {
	// globals that change modes in the microservice framework:
	framework.DefaultSocketPort = 4352 // Global Cache's default socket port
	framework.CheckFunctionAppendBehavior = "Remove older instance"
	framework.RegisterMainGetFunc(doDeviceSpecificGet)
	// framework.RegisterMainSetFunc(doDeviceSpecificSet)
	framework.MicroserviceName = "OpenAV PJLink Microservice"
}

// Every microservice using this golang microservice framework needs to provide this function to invoke functions to do gets.
// socketKey is the network connection for the framework to use to communicate with the device.
// setting is the first parameter in the URI.
// arg1 are the second and third parameters in the URI.
//
//	  Example GET URIs that will result in this function being invoked:
//		 ":address/:setting/"
//	  ":address/:setting/:arg1"
//	  ":address/:setting/:arg1/:arg2"
//
// Every microservice using this golang microservice framework needs to provide this function to invoke functions to do gets.
func doDeviceSpecificGet(socketKey string, setting string, arg1 string, arg2 string) (string, error) {
	function := "doDeviceSpecificGet"

	switch setting {
	case "pinghost":
		return pingHost(socketKey)
	}

	// If we get here, we didn't recognize the setting.  Send an error back to the config writer who had a bad URL.
	errMsg := function + " - unrecognized setting in GET URI: " + setting
	framework.AddToErrors(socketKey, errMsg)
	err := errors.New(errMsg)
	return setting, err
}

func main() {
	setFrameworkGlobals()
	framework.Startup()
}
