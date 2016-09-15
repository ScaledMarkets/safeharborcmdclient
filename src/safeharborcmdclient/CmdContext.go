package main

/*******************************************************************************
 * 
 */
type CmdContext struct {
	RestContext
}

/*******************************************************************************
 * If the specified condition is not true, then print an error message.
 */
func (cmdContext *CmdContext) AssertThat(condition bool, msg string) bool {
	if ! condition {
		cmdContext.FailTest()
		fmt.Println(fmt.Sprintf("ERROR: %s", msg))
	}
	return condition
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) AssertOKResponse(resp *http.Response) {
	if ! cmdContext.Verify200Response(resp) {
		cmdContext.FailTest()
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
	cmdContext.FailTest()
	return false
}

