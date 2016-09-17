package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	//"io/ioutil"
	//"reflect"
	"errors"
	
	"utilities/rest"
)

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetGroupDesc(groupId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetGroupDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getGroupDesc",
		[]string{"Log", "GroupId"},
		[]string{cmdContext.CallDemarcation(), groupId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRepoDesc(repoId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetRepoDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRepoDesc",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.CallDemarcation(), repoId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	rest.PrintMap(responseMap)
	
	// Expect a RepoDesc
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImageDesc(dockerImageId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("getDockerImageDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageDesc",
		[]string{"Log", "DockerImageId"},
		[]string{cmdContext.CallDemarcation(), dockerImageId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) {
		return nil, errors.New(resp.Status)
	}
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	
	// Expect a DockerImageDesc or a DockerImageVersionDesc.
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemDockerfile(dockerfileId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemDockerfile")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remDockerfile",
		[]string{"Log", "DockerfileId"},
		[]string{cmdContext.CallDemarcation(), dockerfileId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerfileDesc(dockerfileId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("getDockerfileDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerfileDesc",
		[]string{"Log", "DockerfileId"},
		[]string{cmdContext.CallDemarcation(), dockerfileId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	
	// Expect a DockerfileDesc
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallCreateRealm(realmName, orgFullName,
	desc string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallCreateRealm")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createRealm",
		[]string{"Log", "RealmName", "OrgFullName", "Description"},
		[]string{cmdContext.CallDemarcation(), realmName, orgFullName, desc})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) TestGetRealmByName(realmName string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("TestGetRealmByName")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmByName",
		[]string{"Log", "RealmName"},
		[]string{cmdContext.CallDemarcation(), realmName})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	// Should return a RealmDesc:
	//	HTTPStatusCode int
	//	HTTPReasonPhrase string
	//	ObjectType string
	//	Id string
	//	RealmName string
	//	OrgFullName string
	//	AdminUserId string
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallCreateUser(userId string, userName string,
	email string, pswd string, realmId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallCreateUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createUser",
		[]string{"Log", "UserId", "UserName", "EmailAddress", "Password", "RealmId"},
		[]string{cmdContext.CallDemarcation(), userId, userName, email, pswd, realmId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallAuthenticate(userId string, pswd string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallAuthenticate")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"authenticate",
		[]string{"Log", "UserId", "Password"},
		[]string{cmdContext.CallDemarcation(), userId, pswd})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDisableUser(userObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallDisableUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"disableUser",
		[]string{"Log", "UserObjId"},
		[]string{cmdContext.CallDemarcation(), userObjId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDeleteGroup(groupId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallDeleteGroup")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"deleteGroup",
		[]string{"Log", "GroupId"},
		[]string{cmdContext.CallDemarcation(), groupId})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallLogout() (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallLogout")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"logout",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallCreateRepo(realmId string, name string,
	desc string, optDockerfilePath string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallCreateRepo")
	
	var resp *http.Response
	var err error
	
	if optDockerfilePath == "" {
		fmt.Println("Using SendSessionPost")
		resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
			"createRepo",
			[]string{"Log", "RealmId", "Name", "Description"},
			[]string{cmdContext.CallDemarcation(), realmId, name, desc})
		fmt.Println("HTTP POST completed")
	} else {
		fmt.Println("Using SendSessionFilePost")
		resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
			"createRepo",
			[]string{"Log", "RealmId", "Name", "Description"},
			[]string{cmdContext.CallDemarcation(), realmId, name, desc},
			optDockerfilePath)
		fmt.Println("HTTP file post completed")
	}
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	// Get the repo Id that is returned in the response body.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallAddDockerfile(repoId string, dockerfilePath string,
	desc string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallAddDockerfile")
	fmt.Println("\t", dockerfilePath)
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"addDockerfile",
		[]string{"Log", "RepoId", "Description"},
		[]string{cmdContext.CallDemarcation(), repoId, desc},
		dockerfilePath)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	// Get the DockerfileDesc that is returned.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerfiles(repoId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetDockerfiles")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerfiles",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.CallDemarcation(), repoId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallExecDockerfile(repoId string, dockerfileId string,
	imageName string, paramNames, paramValues []string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallExecDockerfile")
	
	if len(paramNames) != len(paramValues) { panic(
		"Invalid: len param names != len param values") }
	var paramStr string = ""
	for i, paramName := range paramNames {
		if i > 0 { paramStr = paramStr + ";" }
		paramStr = paramStr + fmt.Sprintf("%s:%s", paramName, paramValues[i])
	}
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"execDockerfile",
		[]string{"Log", "RepoId", "DockerfileId", "ImageName", "Params"},
		[]string{cmdContext.CallDemarcation(), repoId, dockerfileId, imageName, paramStr})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	// Get the repo Id that is returned in the response body.
	/* DockerImageVersionDesc:
	BaseType
	ObjId string
	Version string
	ImageObjId string
    ImageCreationEventId string
    CreationDate string
    Digest []byte
    Signature []byte
    ScanEventIds []string
    DockerBuildOutput string
	*/
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallAddAndExecDockerfile(repoId string, desc string,
	imageName string, dockerfilePath string, paramNames, paramValues []string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallAddAndExecDockerfile")
	
	if len(paramNames) != len(paramValues) { panic(
		"Invalid test: len param names != len param values") }
	var paramStr string = ""
	for i, paramName := range paramNames {
		if i > 0 { paramStr = paramStr + ";" }
		paramStr = paramStr + fmt.Sprintf("%s:%s", paramName, paramValues[i])
	}
	
	var resp *http.Response
	var err error
	//resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
	resp, err = cmdContext.SendSessionFilePost("",
		"addAndExecDockerfile",
		[]string{"Log", "RepoId", "Description", "ImageName", "SessionId", "Params"},
		[]string{cmdContext.CallDemarcation(), repoId, desc, imageName,
			cmdContext.SessionId, paramStr},
		dockerfilePath)

	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	// Returns a DockerImageVersionDesc.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetEventDesc(eventId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetEventDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getEventDesc",
		[]string{"Log", "EventId"},
		[]string{cmdContext.CallDemarcation(), eventId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImages(repoId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetDockerImages")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImages",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.CallDemarcation(), repoId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetUserDesc(userId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetUserDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getUserDesc",
		[]string{"Log", "UserId"},
		[]string{cmdContext.CallDemarcation(), userId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallCreateGroup(realmId, name, description string,
	addMe bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallCreateGroup")
	
	var addMeStr = "false"
	if addMe { addMeStr = "true" }
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createGroup",
		[]string{"Log", "RealmId", "Name", "Description", "AddMe"},
		[]string{cmdContext.CallDemarcation(), realmId, name, description, addMeStr})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err } // returns GroupDesc
	// Id
	// Name
	// Description
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetGroupUsers(groupId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetGroupUsers")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getGroupUsers",
		[]string{"Log", "GroupId"},
		[]string{cmdContext.CallDemarcation(), groupId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)  // returns [UserDesc]
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallAddGroupUser(groupId, userId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallAddGroupUser")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"addGroupUser",
		[]string{"Log", "GroupId", "UserObjId"},
		[]string{cmdContext.CallDemarcation(), groupId, userId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }  // returns Result
	// Status - A value of “0” indicates success.
	// Message
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallMoveUserToRealm(userObjId, realmId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallMoveUserToRealm")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"moveUserToRealm",
		[]string{"Log", "UserObjId", "RealmId"},
		[]string{cmdContext.CallDemarcation(), userObjId, realmId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRealmGroups(realmId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetRealmGroups")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmGroups",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.CallDemarcation(), realmId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)  // returns [GroupDesc]
	if err != nil { fmt.Println(err.Error()); return nil, err }
	// Id
	// Name
	// Description
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRealmRepos(realmId string) (
	[]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetRealmRepos")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmRepos",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.CallDemarcation(), realmId})
	
	if ! cmdContext.Verify200Response(resp) {
		return nil, errors.New(resp.Status)
	}
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	defer resp.Body.Close()
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetAllRealms() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetAllRealms")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getAllRealms",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyDockerfiles() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyDockerfiles")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyDockerfiles",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyDockerImages() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyDockerImages")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyDockerImages",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRealmUsers(realmId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetRealmUsers")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmUsers",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.CallDemarcation(), realmId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallCreateRealmAnon(realmName, orgFullName, adminUserId,
	adminUserFullName, adminEmailAddr, adminPassword string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallCreateRealmAnon")
	
	var resp1 *http.Response
	var err error
	resp1, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createRealmAnon",
		[]string{"Log", "UserId", "UserName", "EmailAddress", "Password", "RealmName", "OrgFullName"},
		[]string{cmdContext.CallDemarcation(), adminUserId, adminUserFullName, adminEmailAddr, adminPassword,
			realmName, orgFullName})
	
		// Returns UserDesc, which contains:
		// Id string
		// UserId string
		// UserName string
		// RealmId string
		
	defer resp1.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp1) { return nil, errors.New(resp1.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp1.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRealmByName(realmName string) (map[string]interface{}, error) {

	cmdContext.StartCall("CallGetRealmByName")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmByName",
		[]string{"Log", "RealmName"},
		[]string{cmdContext.CallDemarcation(), realmName})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallSetPermission(partyId, resourceId string,
	permissions []bool) (map[string]interface{}, error) {

	cmdContext.StartCall("CallSetPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"setPermission",
		[]string{"Log", "PartyId", "ResourceId", "CanCreateIn", "CanRead", "CanWrite", "CanExecute", "CanDelete"},
		[]string{cmdContext.CallDemarcation(), partyId, resourceId, BoolToString(permissions[0]),
			BoolToString(permissions[1]), BoolToString(permissions[2]),
			BoolToString(permissions[3]), BoolToString(permissions[4])})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallAddPermission(partyId, resourceId string,
	permissions []bool) (map[string]interface{}, error) {

	cmdContext.StartCall("CallAddPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"addPermission",
		[]string{"Log", "PartyId", "ResourceId", "CanCreateIn", "CanRead", "CanWrite", "CanExecute", "CanDelete"},
		[]string{cmdContext.CallDemarcation(), partyId, resourceId, BoolToString(permissions[0]),
			BoolToString(permissions[1]), BoolToString(permissions[2]),
			BoolToString(permissions[3]), BoolToString(permissions[4])})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetPermission(partyId, resourceId string) (map[string]interface{}, error) {

	cmdContext.StartCall("CallGetPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getPermission",
		[]string{"Log", "PartyId", "ResourceId"},
		[]string{cmdContext.CallDemarcation(), partyId, resourceId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetScanProviders() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetScanProviders")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getScanProviders",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDefineScanConfig(name, desc, repoId, providerName,
	successExpr, successGraphicFilePath string, providerParamNames []string,
	providerParamValues []string) (map[string]interface{}, error) {

	cmdContext.StartCall("CallDefineScanConfig")
	
	var paramNames []string = []string{"Log", "Name", "Description", "RepoId", "ProviderName"}
	var paramValues []string = []string{cmdContext.CallDemarcation(), name, desc, repoId, providerName}
	paramNames = append(paramNames, providerParamNames...)
	paramValues = append(paramValues, providerParamValues...)
	
	var resp *http.Response
	var err error
	if successGraphicFilePath == "" {
		resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
			"defineScanConfig", paramNames, paramValues)
	} else {
		resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
			"defineScanConfig",
			paramNames,
			paramValues,
			successGraphicFilePath)
	}
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUpdateScanConfig(scanConfigId, name, desc, providerName,
	successExpr, successGraphicFilePath string, providerParamNames []string,
	providerParamValues []string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallUpdateScanConfig")
	
	var paramNames []string = []string{"Log", "ScanConfigId", "Name", "Description", "ProviderName"}
	var paramValues []string = []string{cmdContext.CallDemarcation(), scanConfigId, name, desc, providerName}
	paramNames = append(paramNames, providerParamNames...)
	paramValues = append(paramValues, providerParamValues...)
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"updateScanConfig", paramNames, paramValues, successGraphicFilePath)
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallScanImage(scriptId, imageObjId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallScanImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"scanImage",
		[]string{"Log", "ScanConfigId", "ImageObjId"},
		[]string{cmdContext.CallDemarcation(), scriptId, imageObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	rest.PrintMap(responseMap)
	
	var payload []interface{}
	var isType bool
	payload, isType = responseMap["payload"].([]interface{})
	if !cmdContext.AssertThat(isType, "payload is not a []interface{}") {
		return nil, errors.New(resp.Status)
	}
	
	var eltFieldMaps = make([]map[string]interface{}, 0)
	for _, elt := range payload {
		
		var eltFieldMap map[string]interface{}
		eltFieldMap, isType = elt.(map[string]interface{})
		if cmdContext.AssertThat(isType, "element is not a map[string]interface{}") {
			eltFieldMaps = append(eltFieldMaps, eltFieldMap)
		} else {
			return nil, errors.New(resp.Status)
		}
	}
	
	return eltFieldMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyDesc() (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyDesc",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyGroups() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyGroups")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyGroups",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyRealms() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyRealms")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyRealms",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyRepos() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyRepos")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyRepos",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallReplaceDockerfile(dockerfileId, dockerfilePath,
	desc string) (map[string]interface{}, error) {

	cmdContext.StartCall("CallReplaceDockerfile")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"replaceDockerfile",
		[]string{"Log", "DockerfileId", "Description"},
		[]string{cmdContext.CallDemarcation(), dockerfileId, desc},
		dockerfilePath)
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDownloadImage(imageId, filename string) (int64, error) {

	cmdContext.StartCall("CallDownloadImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"downloadImage",
		[]string{"Log", "ImageObjId"},
		[]string{cmdContext.CallDemarcation(), imageId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }

	if ! cmdContext.Verify200Response(resp) { return 0, errors.New(resp.Status) }
	
	// Check that the server actual sent compressed data
	var reader io.ReadCloser = resp.Body
	var file *os.File
	file, err = os.Create(filename)
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	_, err = io.Copy(file, reader)
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	cmdContext.AssertThat(fileInfo.Size() > 0, "File has zero size")
	fmt.Println("File downloaded to " + filename)
	
	return fileInfo.Size(), nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemGroupUser(groupId, userObjId string) (map[string]interface{}, error) {

	cmdContext.StartCall("CallRemGroupUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remGroupUser",
		[]string{"Log", "GroupId", "UserObjId"},
		[]string{cmdContext.CallDemarcation(), groupId, userObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallReenableUser(userObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallReenableUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"reenableUser",
		[]string{"Log", "UserObjId"},
		[]string{cmdContext.CallDemarcation(), userObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemRealmUser(realmId, userObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemRealmUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remRealmUser",
		[]string{"Log", "RealmId", "UserObjId"},
		[]string{cmdContext.CallDemarcation(), realmId, userObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDeactivateRealm(realmId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallDeactivateRealm")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"deactivateRealm",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.CallDemarcation(), realmId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDeleteRepo(repoId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallDeleteRepo")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"deleteRepo",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.CallDemarcation(), repoId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemPermission(partyId, resourceId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remPermission",
		[]string{"Log", "PartyId", "ResourceId"},
		[]string{cmdContext.CallDemarcation(), partyId, resourceId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetUserEvents(userId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetUserEvents")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getUserEvents",
		[]string{"Log", "UserId"},
		[]string{cmdContext.CallDemarcation(), userId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImageEvents(imageObjId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetDockerImageEvents")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageEvents",
		[]string{"Log", "ImageObjId"},
		[]string{cmdContext.CallDemarcation(), imageObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImageStatus(imageObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetImageStatus")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageStatus",
		[]string{"Log", "ImageObjId"},
		[]string{cmdContext.CallDemarcation(), imageObjId},
		)
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	//EventId string
	//When time.Time
	//UserObjId string
	//EventDescBase
	//ScanConfigId string
	//ProviderName string
    //ParameterValueDescs []*ParameterValueDesc
	//Score string

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerfileEvents(dockerfileId string,
	dockerfilePath string) ([]map[string]interface{}, error) {

	cmdContext.StartCall("CallGetDockerfileEvents")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerfileEvents",
		[]string{"Log", "DockerfileId"},
		[]string{cmdContext.CallDemarcation(), dockerfileId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDefineFlag(repoId, flagName, desc,
	imageFilePath string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallDefineFlag")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"defineFlag",
		[]string{"Log", "RepoId", "Name", "Description"},
		[]string{cmdContext.CallDemarcation(), repoId, flagName, desc},
		imageFilePath)
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetScanConfigDesc(scanConfigId string,
	expectToFindIt bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetScanConfigDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getScanConfigDesc",
		[]string{"Log", "ScanConfigId"},
		[]string{cmdContext.CallDemarcation(), scanConfigId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallChangePassword(userId, oldPswd, newPswd string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallChangePassword")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"changePassword",
		[]string{"Log", "UserId", "OldPassword", "NewPassword"},
		[]string{cmdContext.CallDemarcation(), userId, oldPswd, newPswd})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetFlagDesc(flagId string, expectToFindIt bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetFlagDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getFlagDesc",
		[]string{"Log", "FlagId"},
		[]string{cmdContext.CallDemarcation(), flagId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }

	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetFlagImage(flagId string, filename string) (int64, error) {
	
	cmdContext.StartCall("CallGetFlagImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getFlagImage",
		[]string{"Log", "FlagId"},
		[]string{cmdContext.CallDemarcation(), flagId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	
	if ! cmdContext.Verify200Response(resp) { return 0, errors.New(resp.Status) }
	
	var reader io.ReadCloser = resp.Body
	var file *os.File
	file, err = os.Create(filename)
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	_, err = io.Copy(file, reader)
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if ! cmdContext.AssertErrIsNil(err, "") { return 0, err }
	cmdContext.AssertThat(fileInfo.Size() > 0, "File has zero size")
	fmt.Println("Downloaded flag graphic image " + filename)
	
	return fileInfo.Size(), nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyScanConfigs() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyScanConfigs")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyScanConfigs",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetScanConfigDescByName(repoId, scanConfigName string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetScanConfigDescByName")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getScanConfigDescByName",
		[]string{"Log", "RepoId", "ScanConfigName"},
		[]string{cmdContext.CallDemarcation(), repoId, scanConfigName})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemScanConfig(scanConfigId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemScanConfig")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remScanConfig",
		[]string{"Log", "ScanConfigId"},
		[]string{cmdContext.CallDemarcation(), scanConfigId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) {
		return nil, errors.New(resp.Status)
	}
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyFlags() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetMyFlags")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyFlags",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "while performing SendSessionPost") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetFlagDescByName(repoId, flagName string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetFlagDescByName")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getFlagDescByName",
		[]string{"Log", "RepoId", "FlagName"},
		[]string{cmdContext.CallDemarcation(), repoId, flagName})
	defer resp.Body.Close()
	if err != nil { fmt.Println(err.Error()); return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemFlag(flagId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemFlag")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remFlag",
		[]string{"Log", "FlagId"},
		[]string{cmdContext.CallDemarcation(), flagId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemDockerImage(imageId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemDockerImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remDockerImage",
		[]string{"Log", "ImageId"},
		[]string{cmdContext.CallDemarcation(), imageId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemImageVersion(imageVersionId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallRemImageVersion")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remImageVersion",
		[]string{"Log", "ImageVersionId"},
		[]string{cmdContext.CallDemarcation(), imageVersionId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImageVersions(imageId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("CallGetDockerImageVersions")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageVersions",
		[]string{"Log", "DockerImageId"},
		[]string{cmdContext.CallDemarcation(), imageId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) {
		return nil, errors.New(resp.Status)
	}

	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, err }

	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUpdateUserInfo(userId, userName,
	email string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallUpdateUserInfo")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"updateUserInfo",
		[]string{"Log", "UserId", "UserName", "EmailAddress"},
		[]string{cmdContext.CallDemarcation(), userId, userName, email})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	if ! cmdContext.AssertErrIsNil(err, "when calling updateUserInfo") { return nil, err }
	
	var responseMap map[string]interface{}
	cmdContext.Verify200Response(resp)
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUserExists(userId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallUserExists")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"userExists",
		[]string{"Log", "UserId"},
		[]string{cmdContext.CallDemarcation(), userId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	var responseMap map[string]interface{}
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUseScanConfigForImage(dockerImageId,
	scanConfigId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallUseScanConfigForImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"useScanConfigForImage",
		[]string{"Log", "DockerImageId", "ScanConfigId"},
		[]string{cmdContext.CallDemarcation(), dockerImageId, scanConfigId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallStopUsingScanConfigForImage(dockerImageId,
	scanConfigId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallStopUsingScanConfigForImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"stopUsingScanConfigForImage",
		[]string{"Log", "DockerImageId", "ScanConfigId"},
		[]string{cmdContext.CallDemarcation(), dockerImageId, scanConfigId})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallEnableEmailVerification(enabled bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallEnableEmailVerification")
	
	var resp *http.Response
	var err error
	var flag string
	if enabled {
		flag = "true"
	} else {
		flag = "false"
	}
	
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"enableEmailVerification",
		[]string{"Log", "VerificationEnabled"},
		[]string{cmdContext.CallDemarcation(), flag})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil, err }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallValidateAccountVerificationToken(token string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallValidateAccountVerificationToken")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"validateAccountVerificationToken",
		[]string{"Log", "AccountVerificationToken"},
		[]string{cmdContext.CallDemarcation(), token})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	var responseMap map[string]interface{}
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil, err }
	
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallClearAll() (map[string]interface{}, error) {
	
	cmdContext.StartCall("CallClearAll")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionGet("",
		"clearAll",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, err }
	
	if ! cmdContext.Verify200Response(resp) { return nil, errors.New(resp.Status) }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil, err }
	return responseMap, nil
}
