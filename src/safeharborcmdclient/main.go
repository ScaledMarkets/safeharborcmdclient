/*******************************************************************************
 * Access a SafeHarbor server from a command line.
 * Syntax:
 *	safeharbor [-option]* command [arg]*
 * The commands are:
 *	scan <safeharbor-image-path>
 *	scan from <image-source>
 *	rest <rest-command>
 * where,
 *	<safeharbor-image-path> is the fully qualified name of an image in a
 *		SafeHarbor repository, with the version optional. The syntax is,
 *		<account>/<project>/<image-name>[:<version>]
 *	<image-source> is one of,
 *		docker <image-name>
 *		registry <registry-dns-name>/<image-name>
 *		file <image-file-path>
 *	<rest-command> is the name of the REST function, with arguments according
 *		to that REST function. If a REST function has a file argument, the
 *		command line takes a file path for that argument.
 *
 * Copyright Scaled Markets, Inc.
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"flag"
	"reflect"

	"utilities/rest"
)

const (
	
)

func main() {
	
	var help *bool = flag.Bool("help", false, "Provide help instructions.")
	var scheme *string = flag.String("s", "http", "Protocol scheme (one of http, https, unix)")
	var hostname *string = flag.String("h", "localhost", "Internet address of server.")
	var port *int = flag.Int("p", 80, "Port server is on.")
	var userId *string = flag.String("u", "", "User Id for accessing the Safe Harbor server")
	var password *string = flag.String("w", "", "Password for accessing the Safe Harbor server")
	
	flag.Parse()

	// Create the object on which to make the method call.
	var cmdContext *CmdContext
	cmdContext = CreateCmdContext(*scheme, *hostname, *port,
		*userId, *password, SetSessionId)
	var cmdContextValue = reflect.ValueOf(cmdContext)
	
	// Check arguments.
	if *help {
		fmt.Println("Help:")
		usage(cmdContextValue)
		os.Exit(0)
	}
	
	var args []string = flag.Args()
	if len(args) == 0 {
		usage(cmdContextValue)
		os.Exit(2)
	}
	
	// Obtain the command name and its arguments.
	var command = args[0]
	var commandArgs []string = make([]string, len(args)-1)
	for i, ra := range args {
		if i > 0 {
			commandArgs[i-1] = ra
		}
	}
	
	// Identify the method of CommandContext, using reflection.
	var method = reflect.ValueOf(cmdContext).MethodByName(command)
	if (! method.IsValid()) || method.IsNil() {
		fmt.Println("Method unknown: " + command)
		os.Exit(2)
	}
	if ! method.IsValid() {
		fmt.Println("Method invalid: " + command)
		os.Exit(2)
	}
	
	// Authenticate to server - this returns a SessionId.
	var restResponse map[string]interface{}
	var err error
	if *userId != "" {
		restResponse, err = cmdContext.Authenticate(*userId, *password)
		if err != nil {
			fmt.Println("Authentication with server failed: " + err.Error())
			os.Exit(2)
		}
	
		var obj = restResponse["UniqueSessionId"]
		if obj == nil {
			fmt.Println("Error: UniqueSessionId not found in response from server")
			os.Exit(2)
		}
		var sessionId string
		var isType bool
		sessionId, isType = obj.(string)
		if ! isType {
			fmt.Println("Error: UniqueSessionId is not a string")
			os.Exit(2)
		}
		cmdContext.SetSessionId(sessionId)
	}
	
	// Identify the method to be called.
	var meth reflect.Value = cmdContextValue.MethodByName(command)
	if (! meth.IsValid()) || meth.IsNil() {
		fmt.Println("Method '" + command + "' not found")
		os.Exit(2)
	}

	// Prepare the method parameter values for the method call. These are the
	// command line arguments.
	var inVals = make([]reflect.Value, len(commandArgs))
	//inVals[0] = reflect.ValueOf(cmdContext)
	for i, arg := range commandArgs {
		inVals[i] = reflect.ValueOf(arg)
	}
	
	// Verify that the right number of arguments have been supplied.
	var methodType = meth.Type()
	var numArgs = methodType.NumIn()
	if len(inVals) != numArgs {
		fmt.Println(fmt.Sprintf("Method %s requires %d arguments: %d were supplied",
			command, numArgs, len(inVals)))
		os.Exit(2)
	}
	
	// Perform the method call.
	// All methods return a pair of objects of one of these sets of object types:
	//	map[string]interface{}, error
	//	[]map[string]interface{}, error
	//	int64, error - when a file is downloaded
	var results []reflect.Value
	results = meth.Call(inVals)
	if len(results) != 2 {
		fmt.Println(fmt.Sprintf(
			"%d return value(s) when calling %s: expected two", len(results), command))
		os.Exit(2)
	}
	
	if results[1].IsNil() { // no error was returned
		if results[0].IsNil() {  // error - no result was returned
			
		} else { // ok - there was a result, and no error was returned
			
			// Determine result type.
			switch result := results[0].Interface().(type) {
				case map[string]interface{}:
					rest.PrintMap(result)
				case []map[string]interface{}:
					rest.PrintMaps(result)
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
		var iResult interface{} = results[1]
		err, isType = iResult.(error)
		if ! isType { // error - unexpected type
			fmt.Println("Second return object is not nil but is not an error object")
			os.Exit(2)
		} else { // a valid error object was returned
			fmt.Println("Error: " + err.Error())
			os.Exit(1)
		}
	}
}

/*******************************************************************************
 * 
 */
func SetSessionId(req *http.Request, sessionId string) {
	
	// Set cookie containing the session Id.
	var cookie = &http.Cookie{
		Name: "SessionId",
		Value: sessionId,
		//Path: 
		//Domain: 
		//Expires: 
		//RawExpires: 
		MaxAge: 86400,  // 24 hrs
		Secure: false,  //....change to true later.
		HttpOnly: true,
		//Raw: 
		//Unparsed: 
	}
	
	req.AddCookie(cookie)
}

/*******************************************************************************
 * 
 */
func usage(v reflect.Value) {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [arg]...\n", os.Args[0])
	flag.PrintDefaults()
	
	/*
	fmt.Println("\tCommands:")
	for m := 1; m <= v.NumMethod(); m++ {
		var meth reflect.Value = v.Method(m-1)
		var methodType = meth.Type()
		fmt.Print("\t\t" + methodType.Name())
		var numArgs = methodType.NumIn()
		for a := 1; a <= numArgs; a++ {
			var argType reflect.Type = methodType.In(a-1)
			fmt.Print(" <" + argType.Name() + ">")
		}
		fmt.Println()
	}
	*/
}

/*******************************************************************************
 * 
 */
func BoolToString(b bool) string {
	return fmt.Sprintf("%t", b)
}
