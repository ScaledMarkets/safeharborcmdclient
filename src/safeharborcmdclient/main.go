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
		....
	}
	var command = args[0]
	var restArgs []string = make([]string, len(args))
	restArgs[0] = "Log"
	for i, ra := range args {
		if i > 0 {
			restArgs[i] = ra
		}
	}
	
	// Prepare context.
	var cmdContext = &CmdContext{
		RestContext: utils.NewRestContext(*scheme, *hostname, *port, utils.SetSessionId,
			*stopOnFirstError, *redisPswd, *nolargefiles)
	}
	cmdContext.Print()
	
	// Call the command - a method of CommandContext - using reflection.
	var method = reflect.ValueOf(cmdContext).MethodByName(command)
	if err != nil { ....error }
	if ! method.IsValid() { ...."Method " + command + " is unknown") }
	var v = reflect.ValueOf(....)
	var inVals = make([]reflect.Value, len(restArgs))
	for i, arg := range restArgs {
		inVals[i] = reflect.ValueOf(arg)
	}
	var results []reflect.Value = v.Call(inVals)
	if len(results) != 2 {
		....
	}
	
	var err error
	if results[1].IsNil() {
		var isType bool
		err, isType = results[1].(error)
		if ! isType ....error - unexpected type
	} else {
		if results[0].IsNil() {
			....error - no result
		} else {
			var j []byte
			j, err = json.Marshall(results[0])
			if err != nil {
				....
			} else {
				fmt.Println(string(j))
			}
		}
	}
}
