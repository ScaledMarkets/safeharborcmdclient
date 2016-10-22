package main

import (
	"fmt"
	"net/http"
	"utilities/rest"
)

/*******************************************************************************
 * 
 */
type CmdContext struct {
	rest.RestContext
	SessionId string
	CurrentMethodName string
}

/*******************************************************************************
 * 
 */
func CreateCmdContext(scheme, hostname string, port int, userId string, password string,
	sessionIdSetter func(*http.Request, string)) *CmdContext {
	
	return &CmdContext{
		RestContext: *rest.CreateTCPRestContext(
			scheme, 
			hostname, 
			port, 
			userId, 
			password, 
			sessionIdSetter),
		SessionId: "",
	}
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) SetSessionId(sessionId string) {
	cmdContext.SessionId = sessionId
}

/*******************************************************************************
 * If the specified condition is not true, then print an error message.
 */
func (cmdContext *CmdContext) AssertThat(condition bool, msg string) bool {
	if ! condition {
		fmt.Println(fmt.Sprintf("ERROR: %s", msg))
	}
	return condition
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) AssertOKResponse(resp *http.Response) {
	if ! cmdContext.Verify200Response(resp) {
		fmt.Println("Response status: " + resp.Status)
	}
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) AssertErrIsNil(err error, msg string) bool {
	if err == nil { return true }
	fmt.Println("Original error message:", err.Error())
	fmt.Println("Supplemental message:", msg)
	return false
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) StartCall(methodName string) {
	cmdContext.CurrentMethodName = methodName
	fmt.Println("Calling " + methodName + "...")
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) GetCurrentMethodName() string {
	return cmdContext.CurrentMethodName
}

/*******************************************************************************
 * Write this line to the server''s stdout at the start of each call.
 */
func (cmdContext *CmdContext) CallDemarcation() string {
	return "\n\n" + cmdContext.GetCurrentMethodName() + "<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<"
}

