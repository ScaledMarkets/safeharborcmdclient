/*******************************************************************************
 * Access a SafeHarbor server from a command line.
 * Syntax:
 *	safeharbor [-option]* command [arg]*
 * The commands is simply the name of the REST function, with arguments according
 * to that REST function. If a REST function has a file argument, the command line
 * takes a file path for that argument.
 * The options are:
 *	-TBD-
 */

package main

import (
	"fmt"
	//"net/http"
	"net/url"
	"os"
	"flag"
	"time"
	"strings"
	"reflect"
	"strconv"
	"encoding/json"
	
	// SafeHarbor packages:
	"utilities/rest"
)

const (
	
)

func main() {
	
	var help *bool = flag.Bool("help", false, "Provide help instructions.")
	var scheme *string = flag.String("s", "http", "Protocol scheme (one of http, https, unix)")
	var hostname *string = flag.String("h", "localhost", "Internet address of server.")
	var port *int = flag.Int("p", 80, "Port server is on.")
	var nolargefiles *bool = flag.Bool("nolarge", false, "Do not perform any large file transfers")
	var stopOnFirstError *bool = flag.Bool("stop", false, "Stop after the first error.")
	var redisPswd *string = flag.String("redispswd", "ahdal8934k383898&*kdu&^", "Redis password")
	
	var keys []reflect.Value = reflect.ValueOf(testSuite).MapKeys()
	var allTestNames string
	for i, key := range keys {
		if i > 0 { allTestNames = allTestNames + "," }
		allTestNames = allTestNames + key.String()
	}
	var tests *string = flag.String("tests", allTestNames,
		"Perform the tests listed, comma-separated.")

	flag.Parse()

	if *help {
		fmt.Println("Help:")
		utils.Usage()
		os.Exit(0)
	}
	
	var args []string = flag.Args()
	if len(args) == 0 {
		usage()
		os.Exit(2)
	}
	
	// Obtain the command name and its arguments.
	var command = args[0]
	var commandArgs []string = make([]string, len(args)-1)
	commandArgs[0] = "Log"
	for i, ra := range args {
		if i > 0 {
			commandArgs[i] = ra
		}
	}
	
	// Create the object on which to make the method call.
	var cmdContext = &CmdContext{
		RestContext: utils.NewRestContext(*scheme, *hostname, *port, utils.SetSessionId,
			*stopOnFirstError, *redisPswd, *nolargefiles)
	}
	cmdContext.Print()
	
	// Identify the method of CommandContext, using reflection.
	var method = reflect.ValueOf(cmdContext).MethodByName(command)
	if err != nil {
		fmt.Println("Method unknown: " + command + "; " + err.Error())
		os.Exit(2)
	}
	if ! method.IsValid() {
		fmt.Println("Method invalid: " + command + "; " + err.Error())
		os.Exit(2)
	}
	
	// Prepare the method parameter values for the method call. These are the
	// command line arguments.
	var inVals = make([]reflect.Value, len(commandArgs))
	for i, arg := range commandArgs {
		inVals[i] = reflect.ValueOf(arg)
	}
	
	// Perform the method call.
	// All methods return a pair of objects of one of these sets of object types:
	//	map[string]interface{}, error
	//	[]map[string]interface{}, error
	//	int64, error - when a file is downloaded
	var results []reflect.Value = v.Call(inVals)
	if len(results) != 2 {
		fmt.Println(fmt.Sprintf(
			"%d return value(s) when calling %s: expected two", len(results), command))
		os.Exit(2)
	}
	
	var err error
	if results[1].IsNil() { // no error was returned
		if results[0].IsNil() {  // error - no result was returned
			
		} else { // ok - there was a result, and no error was returned
			
			// Determine result type.
			switch result := results[0].(type) {
				case map[string]interface{}, []map[string]interface{}:
					// Convert the result to JSON and print it.
					var jb []byte
					jb, err = json.Marshall(result)
					if err != nil { // error unmarshalling result
						fmt.Println("Error unmarshalling result: " + err.Error())
						os.Exit(2)
					} else {
						fmt.Println(string(jb))  // Print the result.
					}
					
				case int64: {
					fmt.Println(fmt.Sprintf("%d bytes downloaded", result))
				}
				default: {
					fmt.Println(fmt.Sprintf(
						"Unexpected return object of type %T", result))
					os.Exit(2)
				}
			}
		}
	} else { // an error was returned
		var isType bool
		err, isType = results[1].(error)
		if ! isType { // error - unexpected type
			fmt.Println("Second return object is not nil but is not an error object")
			os.Exit(2)
		} else { // a valid error object was returned
			fmt.Println("Error: " + err.Error())
			os.Exit(1)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
	flag.PrintDefaults()
}
