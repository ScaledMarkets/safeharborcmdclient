package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"errors"
	
	"utilities/rest"
)

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) Ping() (map[string]interface{}, error) {
	
	cmdContext.StartCall("Ping")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"ping",
		[]string{"Log"},
		[]string{cmdContext.CallDemarcation()})
	
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
func (cmdContext *CmdContext) GetGroupDesc(groupId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetGroupDesc")
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
func (cmdContext *CmdContext) GetRepoDesc(repoId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetRepoDesc")
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
func (cmdContext *CmdContext) GetDockerImageDesc(dockerImageId string) (map[string]interface{}, error) {
	
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
func (cmdContext *CmdContext) RemDockerfile(dockerfileId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemDockerfile")
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
func (cmdContext *CmdContext) GetDockerfileDesc(dockerfileId string) (map[string]interface{}, error) {
	
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
func (cmdContext *CmdContext) CreateRealm(realmName, orgFullName,
	desc string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CreateRealm")
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
func (cmdContext *CmdContext) CreateUser(userId string, userName string,
	email string, pswd string, realmId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CreateUser")
	
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
func (cmdContext *CmdContext) Authenticate(userId string, pswd string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("Authenticate")
	
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
func (cmdContext *CmdContext) DisableUser(userObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("DisableUser")
	
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
func (cmdContext *CmdContext) DeleteGroup(groupId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("DeleteGroup")
	
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
func (cmdContext *CmdContext) Logout() (map[string]interface{}, error) {
	
	cmdContext.StartCall("Logout")
	
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
func (cmdContext *CmdContext) CreateRepo(realmId string, name string,
	desc string, optDockerfilePath string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CreateRepo")
	
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
func (cmdContext *CmdContext) AddDockerfile(repoId string, dockerfilePath string,
	desc string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("AddDockerfile")
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
func (cmdContext *CmdContext) GetDockerfiles(repoId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetDockerfiles")
	
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
func (cmdContext *CmdContext) ExecDockerfile(repoId string, dockerfileId string,
	imageName string, paramNames, paramValues []string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("ExecDockerfile")
	
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
func (cmdContext *CmdContext) AddAndExecDockerfile(repoId string, desc string,
	imageName string, dockerfilePath string, paramNames, paramValues []string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("AddAndExecDockerfile")
	
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
func (cmdContext *CmdContext) GetEventDesc(eventId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetEventDesc")
	
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
func (cmdContext *CmdContext) GetDockerImages(repoId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetDockerImages")
	
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
func (cmdContext *CmdContext) GetUserDesc(userId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetUserDesc")
	
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
func (cmdContext *CmdContext) CreateGroup(realmId, name, description string,
	addMe bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CreateGroup")
	
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
func (cmdContext *CmdContext) GetGroupUsers(groupId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetGroupUsers")

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
func (cmdContext *CmdContext) AddGroupUser(groupId, userId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("AddGroupUser")

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
func (cmdContext *CmdContext) MoveUserToRealm(userObjId, realmId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("MoveUserToRealm")
	
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
func (cmdContext *CmdContext) GetRealmGroups(realmId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetRealmGroups")

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
func (cmdContext *CmdContext) GetRealmRepos(realmId string) (
	[]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetRealmRepos")
	
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
func (cmdContext *CmdContext) GetAllRealms() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetAllRealms")
	
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
func (cmdContext *CmdContext) GetMyDockerfiles() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyDockerfiles")
	
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
func (cmdContext *CmdContext) GetMyDockerImages() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyDockerImages")
	
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
func (cmdContext *CmdContext) GetRealmUsers(realmId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetRealmUsers")
	
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
func (cmdContext *CmdContext) CreateRealmAnon(realmName, orgFullName, adminUserId,
	adminUserFullName, adminEmailAddr, adminPassword string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("CreateRealmAnon")
	
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
func (cmdContext *CmdContext) GetRealmByName(realmName string) (map[string]interface{}, error) {

	cmdContext.StartCall("GetRealmByName")
	
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
func (cmdContext *CmdContext) SetPermission(partyId, resourceId string,
	permissions []bool) (map[string]interface{}, error) {

	cmdContext.StartCall("SetPermission")
	
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
func (cmdContext *CmdContext) AddPermission(partyId, resourceId string,
	permissions []bool) (map[string]interface{}, error) {

	cmdContext.StartCall("AddPermission")
	
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
func (cmdContext *CmdContext) GetPermission(partyId, resourceId string) (map[string]interface{}, error) {

	cmdContext.StartCall("GetPermission")
	
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
func (cmdContext *CmdContext) GetScanProviders() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetScanProviders")
	
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
func (cmdContext *CmdContext) DefineScanConfig(name, desc, repoId, providerName,
	successExpr, successGraphicFilePath string, providerParamNames []string,
	providerParamValues []string) (map[string]interface{}, error) {

	cmdContext.StartCall("DefineScanConfig")
	
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
func (cmdContext *CmdContext) UpdateScanConfig(scanConfigId, name, desc, providerName,
	successExpr, successGraphicFilePath string, providerParamNames []string,
	providerParamValues []string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("UpdateScanConfig")
	
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
func (cmdContext *CmdContext) ScanImage(scriptId, imageObjId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("ScanImage")
	
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
func (cmdContext *CmdContext) GetMyDesc() (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyDesc")
	
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
func (cmdContext *CmdContext) GetMyGroups() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyGroups")
	
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
func (cmdContext *CmdContext) GetMyRealms() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyRealms")
	
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
func (cmdContext *CmdContext) GetMyRepos() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyRepos")
	
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
func (cmdContext *CmdContext) ReplaceDockerfile(dockerfileId, dockerfilePath,
	desc string) (map[string]interface{}, error) {

	cmdContext.StartCall("ReplaceDockerfile")
	
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
func (cmdContext *CmdContext) DownloadImage(imageId, filename string) (int64, error) {

	cmdContext.StartCall("DownloadImage")
	
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
func (cmdContext *CmdContext) RemGroupUser(groupId, userObjId string) (map[string]interface{}, error) {

	cmdContext.StartCall("RemGroupUser")
	
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
func (cmdContext *CmdContext) ReenableUser(userObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("ReenableUser")
	
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
func (cmdContext *CmdContext) RemRealmUser(realmId, userObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemRealmUser")
	
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
func (cmdContext *CmdContext) DeactivateRealm(realmId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("DeactivateRealm")
	
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
func (cmdContext *CmdContext) DeleteRepo(repoId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("DeleteRepo")
	
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
func (cmdContext *CmdContext) RemPermission(partyId, resourceId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemPermission")
	
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
func (cmdContext *CmdContext) GetUserEvents(userId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetUserEvents")
	
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
func (cmdContext *CmdContext) GetDockerImageEvents(imageObjId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetDockerImageEvents")
	
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
func (cmdContext *CmdContext) GetDockerImageStatus(imageObjId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetImageStatus")
	
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
func (cmdContext *CmdContext) GetDockerfileEvents(dockerfileId string,
	dockerfilePath string) ([]map[string]interface{}, error) {

	cmdContext.StartCall("GetDockerfileEvents")
	
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
func (cmdContext *CmdContext) DefineFlag(repoId, flagName, desc,
	imageFilePath string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("DefineFlag")
	
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
func (cmdContext *CmdContext) GetScanConfigDesc(scanConfigId string,
	expectToFindIt bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetScanConfigDesc")
	
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
func (cmdContext *CmdContext) ChangePassword(userId, oldPswd, newPswd string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("ChangePassword")
	
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
func (cmdContext *CmdContext) GetFlagDesc(flagId string, expectToFindIt bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetFlagDesc")
	
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
func (cmdContext *CmdContext) GetFlagImage(flagId string, filename string) (int64, error) {
	
	cmdContext.StartCall("GetFlagImage")
	
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
func (cmdContext *CmdContext) GetMyScanConfigs() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyScanConfigs")
	
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
func (cmdContext *CmdContext) GetScanConfigDescByName(repoId, scanConfigName string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetScanConfigDescByName")
	
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
func (cmdContext *CmdContext) RemScanConfig(scanConfigId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemScanConfig")
	
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
func (cmdContext *CmdContext) GetMyFlags() ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetMyFlags")
	
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
func (cmdContext *CmdContext) GetFlagDescByName(repoId, flagName string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("GetFlagDescByName")
	
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
func (cmdContext *CmdContext) RemFlag(flagId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemFlag")
	
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
func (cmdContext *CmdContext) RemDockerImage(imageId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemDockerImage")
	
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
func (cmdContext *CmdContext) RemImageVersion(imageVersionId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("RemImageVersion")

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
func (cmdContext *CmdContext) GetDockerImageVersions(imageId string) ([]map[string]interface{}, error) {
	
	cmdContext.StartCall("GetDockerImageVersions")

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
func (cmdContext *CmdContext) UpdateUserInfo(userId, userName,
	email string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("UpdateUserInfo")

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
func (cmdContext *CmdContext) UserExists(userId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("UserExists")

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
func (cmdContext *CmdContext) UseScanConfigForImage(dockerImageId,
	scanConfigId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("UseScanConfigForImage")
	
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
func (cmdContext *CmdContext) StopUsingScanConfigForImage(dockerImageId,
	scanConfigId string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("StopUsingScanConfigForImage")
	
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
func (cmdContext *CmdContext) EnableEmailVerification(enabled bool) (map[string]interface{}, error) {
	
	cmdContext.StartCall("EnableEmailVerification")
	
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
func (cmdContext *CmdContext) ValidateAccountVerificationToken(token string) (map[string]interface{}, error) {
	
	cmdContext.StartCall("ValidateAccountVerificationToken")
	
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
func (cmdContext *CmdContext) ClearAll() (map[string]interface{}, error) {
	
	cmdContext.StartCall("ClearAll")
	
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
