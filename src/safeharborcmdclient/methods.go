package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"io/ioutil"
	"reflect"
	
	"utilities/rest"
)

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetGroupDesc(groupId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetGroupDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getGroupDesc",
		[]string{"Log", "GroupId"},
		[]string{cmdContext.TestDemarcation(), groupId})
	
	defer resp.Body.Close()
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil, err }
	
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRepoDesc(repoId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetRepoDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRepoDesc",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.TestDemarcation(), repoId})
	
	defer resp.Body.Close()
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return }
	rest.PrintMap(responseMap)
	
	// Expect a RepoDesc
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImageDesc(dockerImageId string) (map[string]interface{}, error) {
	
	cmdContext.StartTest("getDockerImageDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageDesc",
		[]string{"Log", "DockerImageId"},
		[]string{cmdContext.TestDemarcation(), dockerImageId})
	
	defer resp.Body.Close()
	
	if ! cmdContext.Verify200Response(resp) {
		cmdContext.FailTest()
		return nil
	}
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil }
	
	// Expect a DockerImageDesc or a DockerImageVersionDesc.
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemDockerfile(dockerfileId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemDockerfile")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remDockerfile",
		[]string{"Log", "DockerfileId"},
		[]string{cmdContext.TestDemarcation(), dockerfileId})
	
	defer resp.Body.Close()
	if err != nil {
		cmdContext.FailTest()
		return
	}

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil }
	
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerfileDesc(dockerfileId string) (map[string]interface{}, error) {
	
	cmdContext.StartTest("getDockerfileDesc")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerfileDesc",
		[]string{"Log", "DockerfileId"},
		[]string{cmdContext.TestDemarcation(), dockerfileId})
	
	defer resp.Body.Close()
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap") { return nil }
	
	// Expect a DockerfileDesc
	return responseMap, nil
}

/*******************************************************************************
 * Verify that we can create a new realm.
 */
func (cmdContext *CmdContext) CallCreateRealm(realmName, orgFullName,
	desc string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallCreateRealm")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createRealm",
		[]string{"Log", "RealmName", "OrgFullName", "Description"},
		[]string{cmdContext.TestDemarcation(), realmName, orgFullName, desc})
	
	defer resp.Body.Close()
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	// Get the realm Id that is returned in the response body.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "" }
	return responseMap, nil
}

/*******************************************************************************
 * Verify that we can look up a realm by its name.
 */
func (cmdContext *CmdContext) TestGetRealmByName(realmName string) (map[string]interface{}, error) {
	
	cmdContext.StartTest("TestGetRealmByName")
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmByName",
		[]string{"Log", "RealmName"},
		[]string{cmdContext.TestDemarcation(), realmName})
	
	defer resp.Body.Close()
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	// Get the realm Id that is returned in the response body.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	
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
 * Return the object Id of the new user.
 */
func (cmdContext *CmdContext) CallCreateUser(userId string, userName string,
	email string, pswd string, realmId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallCreateUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createUser",
		[]string{"Log", "UserId", "UserName", "EmailAddress", "Password", "RealmId"},
		[]string{cmdContext.TestDemarcation(), userId, userName, email, pswd, realmId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "", nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns session Id, and isAdmin.
 */
func (cmdContext *CmdContext) CallAuthenticate(userId string, pswd string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallAuthenticate")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"authenticate",
		[]string{"Log", "UserId", "Password"},
		[]string{cmdContext.TestDemarcation(), userId, pswd})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	// Get the repo Id that is returned in the response body.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "", false }
	return responseMap, nil
}

/*******************************************************************************
 * Return true if successful.
 */
func (cmdContext *CmdContext) CallDisableUser(userObjId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallDisableUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"disableUser",
		[]string{"Log", "UserObjId"},
		[]string{cmdContext.TestDemarcation(), userObjId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return false }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDeleteGroup(groupId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallDeleteGroup")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"deleteGroup",
		[]string{"Log", "GroupId"},
		[]string{cmdContext.TestDemarcation(), groupId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	return responseMap, nil
}

/*******************************************************************************
 * If successful, return true.
 */
func (cmdContext *CmdContext) CallLogout() (map[string]interface{}, error) {
	
	restContext.StartCall("CallLogout")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"logout",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return false }
	return responseMap, nil
}

/*******************************************************************************
 * Verify that we can create a new repo. This requires that we first created
 * a realm that the repo can belong to.
 */
func (cmdContext *CmdContext) CallCreateRepo(realmId string, name string,
	desc string, optDockerfilePath string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallCreateRepo")
	
	var resp *http.Response
	var err error
	
	if optDockerfilePath == "" {
		fmt.Println("Using SendSessionPost")
		resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
			"createRepo",
			[]string{"Log", "RealmId", "Name", "Description"},
			[]string{cmdContext.TestDemarcation(), realmId, name, desc})
		fmt.Println("HTTP POST completed")
	} else {
		fmt.Println("Using SendSessionFilePost")
		resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
			"createRepo",
			[]string{"Log", "RealmId", "Name", "Description"},
			[]string{cmdContext.TestDemarcation(), realmId, name, desc},
			optDockerfilePath)
		fmt.Println("HTTP file post completed")
	}
	if ! cmdContext.AssertErrIsNil(err, "") { return "" }
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	// Get the repo Id that is returned in the response body.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "" }
	return responseMap, nil
}

/*******************************************************************************
 * Verify that we can upload a dockerfile. This requries that we first created
 * a repo to uplaod it into. Returns the Id of the dockerfile, and a map of the
 * fields of the DockerfileDesc.
 */
func (cmdContext *CmdContext) CallAddDockerfile(repoId string, dockerfilePath string,
	desc string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallAddDockerfile")
	fmt.Println("\t", dockerfilePath)
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"addDockerfile",
		[]string{"Log", "RepoId", "Description"},
		[]string{cmdContext.TestDemarcation(), repoId, desc},
		dockerfilePath)
	if ! cmdContext.AssertErrIsNil(err, "") { return "", nil }
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	// Get the DockerfileDesc that is returned.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "", nil }
	return responseMap, nil
}

/*******************************************************************************
 * Verify that we can obtain the names of the dockerfiles owned by the specified
 * repo. The result is an array of dockerfile names.
 */
func (cmdContext *CmdContext) CallGetDockerfiles(repoId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetDockerfiles")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerfiles",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.TestDemarcation(), repoId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * Verify that we can build an image, from a dockerfile that has already been
 * uploaded into a repo and for which we have the SafeHarborServer image id.
 * The result is the object Id of the image version, and the image.
 */
func (cmdContext *CmdContext) CallExecDockerfile(repoId string, dockerfileId string,
	imageName string, paramNames, paramValues []string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallExecDockerfile")
	
	if len(paramNames) != len(paramValues) { panic(
		"Invalid test: len param names != len param values") }
	var paramStr string = ""
	for i, paramName := range paramNames {
		if i > 0 { paramStr = paramStr + ";" }
		paramStr = paramStr + fmt.Sprintf("%s:%s", paramName, paramValues[i])
	}
	
	fmt.Println("paramStr=" + paramStr)
	fmt.Println(fmt.Sprintf("len(paramNames)=%d, len(paramValues)=%d", len(paramNames), len(paramValues)))
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"execDockerfile",
		[]string{"Log", "RepoId", "DockerfileId", "ImageName", "Params"},
		[]string{cmdContext.TestDemarcation(), repoId, dockerfileId, imageName, paramStr})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
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
	if err != nil { fmt.Println(err.Error()); return "", "", nil }
	return responseMap, nil
}

/*******************************************************************************
 * Verify that we can upload a dockerfile and build an image from it.
 * The result is the object Id of the image version, and the image,
 * and the object Id of the event pertaining to the creation of the image.
 */
func (cmdContext *CmdContext) CallAddAndExecDockerfile(repoId string, desc string,
	imageName string, dockerfilePath string, paramNames, paramValues []string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallAddAndExecDockerfile")
	
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
		[]string{cmdContext.TestDemarcation(), repoId, desc, imageName,
			cmdContext.SessionId, paramStr},
		dockerfilePath)

	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	// Returns a DockerImageVersionDesc.
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "", "", "", nil }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetEventDesc(eventId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetEventDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getEventDesc",
		[]string{"Log", "EventId"},
		[]string{cmdContext.TestDemarcation(), eventId})
	defer resp.Body.Close()
	if err != nil {
		cmdContext.FailTest()
		return nil
	}

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Result is an array of the names of the images owned by the specified repo.
 */
func (cmdContext *CmdContext) CallGetDockerImages(repoId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetDockerImages")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImages",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.TestDemarcation(), repoId})
	
	defer resp.Body.Close()
	if err != nil {
		cmdContext.FailTest()
		return nil
	}

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * Return the object Id of the specified user, and a list of the realms that
 * the user can modify.
 */
func (cmdContext *CmdContext) CallGetUserDesc(userId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetUserDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getUserDesc",
		[]string{"Log", "UserId"},
		[]string{cmdContext.TestDemarcation(), userId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallCreateGroup(realmId, name, description string,
	addMe bool) (map[string]interface{}, error) {
	
	restContext.StartCall("CallCreateGroup")
	
	var addMeStr = "false"
	if addMe { addMeStr = "true" }
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createGroup",
		[]string{"Log", "RealmId", "Name", "Description", "AddMe"},
		[]string{cmdContext.TestDemarcation(), realmId, name, description, addMeStr})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "" } // returns GroupDesc
	// Id
	// Name
	// Description
	return responseMap, nil
}

/*******************************************************************************
 * Return an array of the user object ids.
 */
func (cmdContext *CmdContext) CallGetGroupUsers(groupId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetGroupUsers")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getGroupUsers",
		[]string{"Log", "GroupId"},
		[]string{cmdContext.TestDemarcation(), groupId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)  // returns [UserDesc]
	if err != nil { fmt.Println(err.Error()); return nil }
	
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallAddGroupUser(groupId, userId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallAddGroupUser")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"addGroupUser",
		[]string{"Log", "GroupId", "UserObjId"},
		[]string{cmdContext.TestDemarcation(), groupId, userId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return false }  // returns Result
	// Status - A value of “0” indicates success.
	// Message
	return responseMap, nil
}

/*******************************************************************************
 * Returns result.
 */
func (cmdContext *CmdContext) CallMoveUserToRealm(userObjId, realmId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallMoveUserToRealm")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"moveUserToRealm",
		[]string{"Log", "UserObjId", "RealmId"},
		[]string{cmdContext.TestDemarcation(), userObjId, realmId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRealmGroups(realmId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetRealmGroups")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmGroups",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.TestDemarcation(), realmId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)  // returns [GroupDesc]
	if err != nil { fmt.Println(err.Error()); return nil }
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
	
	restContext.StartCall("CallGetRealmRepos")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmRepos",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.TestDemarcation(), realmId})
	
	if ! cmdContext.Verify200Response(resp) {
		cmdContext.FailTest()
	}
	
	defer resp.Body.Close()
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetAllRealms() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetAllRealms")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getAllRealms",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns the Ids of the dockerfiles.
 */
func (cmdContext *CmdContext) CallGetMyDockerfiles() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyDockerfiles")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyDockerfiles",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns the Ids of the image objects.
 */
func (cmdContext *CmdContext) CallGetMyDockerImages() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyDockerImages")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyDockerImages",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * Returns the obj Ids of the realm''s users.
 */
func (cmdContext *CmdContext) CallGetRealmUsers(realmId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetRealmUsers")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmUsers",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.TestDemarcation(), realmId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns the (Id, Id) of the created realm and user, respectively
 */
func (cmdContext *CmdContext) CallCreateRealmAnon(realmName, orgFullName, adminUserId,
	adminUserFullName, adminEmailAddr, adminPassword string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallCreateRealmAnon")
	
	var resp1 *http.Response
	var err error
	resp1, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"createRealmAnon",
		[]string{"Log", "UserId", "UserName", "EmailAddress", "Password", "RealmName", "OrgFullName"},
		[]string{cmdContext.TestDemarcation(), adminUserId, adminUserFullName, adminEmailAddr, adminPassword,
			realmName, orgFullName})
	
		// Returns UserDesc, which contains:
		// Id string
		// UserId string
		// UserName string
		// RealmId string
		
	if err != nil { fmt.Println(err.Error()); return "", "", nil }

	defer resp1.Body.Close()

	if ! cmdContext.Verify200Response(resp1) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp1.Body)
	if err != nil { fmt.Println(err.Error()); return "", "", nil }

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetRealmByName(realmName string) (map[string]interface{}, error) {

	restContext.StartCall("CallGetRealmByName")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getRealmByName",
		[]string{"Log", "RealmName"},
		[]string{cmdContext.TestDemarcation(), realmName})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns the permissions that resulted.
 */
func (cmdContext *CmdContext) CallSetPermission(partyId, resourceId string,
	permissions []bool) (map[string]interface{}, error) {

	restContext.StartCall("CallSetPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"setPermission",
		[]string{"Log", "PartyId", "ResourceId", "CanCreateIn", "CanRead", "CanWrite", "CanExecute", "CanDelete"},
		[]string{cmdContext.TestDemarcation(), partyId, resourceId, BoolToString(permissions[0]),
			BoolToString(permissions[1]), BoolToString(permissions[2]),
			BoolToString(permissions[3]), BoolToString(permissions[4])})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns the permissions that resulted.
 */
func (cmdContext *CmdContext) CallAddPermission(partyId, resourceId string,
	permissions []bool) (map[string]interface{}, error) {

	restContext.StartCall("CallAddPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"addPermission",
		[]string{"Log", "PartyId", "ResourceId", "CanCreateIn", "CanRead", "CanWrite", "CanExecute", "CanDelete"},
		[]string{cmdContext.TestDemarcation(), partyId, resourceId, BoolToString(permissions[0]),
			BoolToString(permissions[1]), BoolToString(permissions[2]),
			BoolToString(permissions[3]), BoolToString(permissions[4])})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Return an array of string representing the values for the permission mask.
 */
func (cmdContext *CmdContext) CallGetPermission(partyId, resourceId string) (map[string]interface{}, error) {

	restContext.StartCall("CallGetPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getPermission",
		[]string{"Log", "PartyId", "ResourceId"},
		[]string{cmdContext.TestDemarcation(), partyId, resourceId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Return an array of the names of the available providers.
 */
func (cmdContext *CmdContext) CallGetScanProviders() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetScanProviders")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getScanProviders",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return }
	return responseMaps, nil
}

/*******************************************************************************
 * Returns the Id of the ScanConfig that gets created.
 */
func (cmdContext *CmdContext) CallDefineScanConfig(name, desc, repoId, providerName,
	successExpr, successGraphicFilePath string, providerParamNames []string,
	providerParamValues []string) (map[string]interface{}, error) {

	restContext.StartCall("CallDefineScanConfig")
	
	var paramNames []string = []string{"Log", "Name", "Description", "RepoId", "ProviderName"}
	var paramValues []string = []string{cmdContext.TestDemarcation(), name, desc, repoId, providerName}
	paramNames = append(paramNames, providerParamNames...)
	paramValues = append(paramValues, providerParamValues...)
	
	fmt.Println("Param names:")
	for _, n := range paramNames { fmt.Println("\t" + n) }
	fmt.Println("Param values:")
	for _, v := range paramValues { fmt.Println("\t" + v) }
	
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
	cmdContext.AssertErrIsNil(err, "at the POST")
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	cmdContext.AssertErrIsNil(err, "at ParseResponseBodyToMap")
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUpdateScanConfig(scanConfigId, name, desc, providerName,
	successExpr, successGraphicFilePath string, providerParamNames []string,
	providerParamValues []string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallUpdateScanConfig")
	
	var paramNames []string = []string{"Log", "ScanConfigId", "Name", "Description", "ProviderName"}
	var paramValues []string = []string{cmdContext.TestDemarcation(), scanConfigId, name, desc, providerName}
	paramNames = append(paramNames, providerParamNames...)
	paramValues = append(paramValues, providerParamValues...)
	
	fmt.Println("Param names:")
	for _, n := range paramNames { fmt.Println("\t" + n) }
	fmt.Println("Param values:")
	for _, v := range paramValues { fmt.Println("\t" + v) }
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"updateScanConfig", paramNames, paramValues, successGraphicFilePath)
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil }
	return responseMap, nil
}

/*******************************************************************************
 * Returns array of maps, each containing the fields of a ScanEventDesc.
 */
func (cmdContext *CmdContext) CallScanImage(scriptId, imageObjId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallScanImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"scanImage",
		[]string{"Log", "ScanConfigId", "ImageObjId"},
		[]string{cmdContext.TestDemarcation(), scriptId, imageObjId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	rest.PrintMap(responseMap)
	
	var payload []interface{}
	var isType bool
	payload, isType = responseMap["payload"].([]interface{})
	if !cmdContext.AssertThat(isType, "payload is not a []interface{}") {
		cmdContext.FailTest()
		return nil
	}
	
	var eltFieldMaps = make([]map[string]interface{}, 0)
	for _, elt := range payload {
		
		var eltFieldMap map[string]interface{}
		eltFieldMap, isType = elt.(map[string]interface{})
		if cmdContext.AssertThat(isType, "element is not a map[string]interface{}") {
			eltFieldMaps = append(eltFieldMaps, eltFieldMap)
		} else {
			cmdContext.FailTest()
			return nil
		}
	}
	
	return eltFieldMaps, nil
}

/*******************************************************************************
 * Return the object Id of the current authenticated user.
 */
func (cmdContext *CmdContext) CallGetMyDesc() (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyDesc",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if err != nil { fmt.Println(err.Error()); return "", nil }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyGroups() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyGroups")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyGroups",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyRealms() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyRealms")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyRealms",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyRepos() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyRepos")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyRepos",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallReplaceDockerfile(dockerfileId, dockerfilePath,
	desc string) (map[string]interface{}, error) {

	restContext.StartCall("CallReplaceDockerfile")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"replaceDockerfile",
		[]string{"Log", "DockerfileId", "Description"},
		[]string{cmdContext.TestDemarcation(), dockerfileId, desc},
		dockerfilePath)
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDownloadImage(imageId, filename string) (map[string]interface{}, error) {

	restContext.StartCall("CallDownloadImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"downloadImage",
		[]string{"Log", "ImageObjId"},
		[]string{cmdContext.TestDemarcation(), imageId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	// Check that the server actual sent compressed data
	var reader io.ReadCloser = resp.Body
	var file *os.File
	file, err = os.Create(filename)
	cmdContext.AssertErrIsNil(err, "")
	_, err = io.Copy(file, reader)
	cmdContext.AssertErrIsNil(err, "")
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	cmdContext.AssertThat(fileInfo.Size() > 0, "File has zero size")
	fmt.Println("File downloaded to " + filename)
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemGroupUser(groupId, userObjId string) (map[string]interface{}, error) {

	restContext.StartCall("CallRemGroupUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remGroupUser",
		[]string{"Log", "GroupId", "UserObjId"},
		[]string{cmdContext.TestDemarcation(), groupId, userObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallReenableUser(userObjId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallReenableUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"reenableUser",
		[]string{"Log", "UserObjId"},
		[]string{cmdContext.TestDemarcation(), userObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	return responseMap, nil
}
	
/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemRealmUser(realmId, userObjId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemRealmUser")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remRealmUser",
		[]string{"Log", "RealmId", "UserObjId"},
		[]string{cmdContext.TestDemarcation(), realmId, userObjId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDeactivateRealm(realmId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallDeactivateRealm")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"deactivateRealm",
		[]string{"Log", "RealmId"},
		[]string{cmdContext.TestDemarcation(), realmId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallDeleteRepo(repoId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallDeleteRepo")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"deleteRepo",
		[]string{"Log", "RepoId"},
		[]string{cmdContext.TestDemarcation(), repoId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemPermission(partyId, resourceId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemPermission")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remPermission",
		[]string{"Log", "PartyId", "ResourceId"},
		[]string{cmdContext.TestDemarcation(), partyId, resourceId})
	
	defer resp.Body.Close()
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetUserEvents(userId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetUserEvents")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getUserEvents",
		[]string{"Log", "UserId"},
		[]string{cmdContext.TestDemarcation(), userId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * Returns array of event Ids.
 */
func (cmdContext *CmdContext) CallGetDockerImageEvents(imageObjId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetDockerImageEvents")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageEvents",
		[]string{"Log", "ImageObjId"},
		[]string{cmdContext.TestDemarcation(), imageObjId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetDockerImageStatus(imageObjId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetImageStatus")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageStatus",
		[]string{"Log", "ImageObjId"},
		[]string{cmdContext.TestDemarcation(), imageObjId},
		)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil }

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

	restContext.StartCall("CallGetDockerfileEvents")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerfileEvents",
		[]string{"Log", "DockerfileId"},
		[]string{cmdContext.TestDemarcation(), dockerfileId})
	
	defer resp.Body.Close()

	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
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
	
	restContext.StartCall("CallDefineFlag")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionFilePost(cmdContext.SessionId,
		"defineFlag",
		[]string{"Log", "RepoId", "Name", "Description"},
		[]string{cmdContext.TestDemarcation(), repoId, flagName, desc},
		imageFilePath)
	if ! cmdContext.AssertErrIsNil(err, "") { return nil}
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	cmdContext.AssertErrIsNil(err, "")

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetScanConfigDesc(scanConfigId string,
	expectToFindIt bool) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetScanConfigDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getScanConfigDesc",
		[]string{"Log", "ScanConfigId"},
		[]string{cmdContext.TestDemarcation(), scanConfigId})
	if ! cmdContext.AssertErrIsNil(err, "") { return nil }
	
	if expectToFindIt {
		if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	} else {
		if resp.StatusCode == 200 {
			cmdContext.FailTest()
		} else {
		}	
		return nil
	}
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	cmdContext.AssertErrIsNil(err, "")
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallChangePassword(userId, oldPswd, newPswd string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallChangePassword")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"changePassword",
		[]string{"Log", "UserId", "OldPassword", "NewPassword"},
		[]string{cmdContext.TestDemarcation(), userId, oldPswd, newPswd})
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	cmdContext.AssertErrIsNil(err, "")

	return responseMap, nil
}

/*******************************************************************************
 * Returns the name of the flag.
 */
func (cmdContext *CmdContext) CallGetFlagDesc(flagId string, expectToFindIt bool) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetFlagDesc")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getFlagDesc",
		[]string{"Log", "FlagId"},
		[]string{cmdContext.TestDemarcation(), flagId})
	if ! cmdContext.AssertErrIsNil(err, "") { return ""}
	
	if expectToFindIt {
		if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	} else {
		if resp.StatusCode == 200 {
			cmdContext.FailTest()
		} else {
		}	
		return ""
	}

	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return "" }
	return responseMap, nil
}

/*******************************************************************************
 * Returns the size of the file that was downloaded.
 */
func (cmdContext *CmdContext) CallGetFlagImage(flagId string, filename string) (int64, error) {
	
	restContext.StartCall("CallGetFlagImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getFlagImage",
		[]string{"Log", "FlagId"},
		[]string{cmdContext.TestDemarcation(), flagId})
	if ! cmdContext.AssertErrIsNil(err, "") { return 0 }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var reader io.ReadCloser = resp.Body
	var file *os.File
	file, err = os.Create(filename)
	cmdContext.AssertErrIsNil(err, "")
	_, err = io.Copy(file, reader)
	cmdContext.AssertErrIsNil(err, "")
	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if ! cmdContext.AssertErrIsNil(err, "") { return 0 }
	cmdContext.AssertThat(fileInfo.Size() > 0, "File has zero size")
	fmt.Println("Downloaded flag graphic image " + filename)
	
	return fileInfo.Size(), nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyScanConfigs() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyScanConfigs")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyScanConfigs",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	if ! cmdContext.AssertErrIsNil(err, "") { return nil, nil }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil, nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetScanConfigDescByName(repoId, scanConfigName string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetScanConfigDescByName")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getScanConfigDescByName",
		[]string{"Log", "RepoId", "ScanConfigName"},
		[]string{cmdContext.TestDemarcation(), repoId, scanConfigName})
	if ! cmdContext.AssertErrIsNil(err, "") { return "" }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return "" }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemScanConfig(scanConfigId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemScanConfig")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remScanConfig",
		[]string{"Log", "ScanConfigId"},
		[]string{cmdContext.TestDemarcation(), scanConfigId})
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	
	if ! cmdContext.Verify200Response(resp) {
		cmdContext.FailTest()
		return false
	}
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return false }

	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetMyFlags() ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetMyFlags")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getMyFlags",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	if ! cmdContext.AssertErrIsNil(err, "while performing SendSessionPost") { return nil }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }
	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallGetFlagDescByName(repoId, flagName string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallGetFlagDescByName")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getFlagDescByName",
		[]string{"Log", "RepoId", "FlagName"},
		[]string{cmdContext.TestDemarcation(), repoId, flagName})
	if err != nil { fmt.Println(err.Error()); return "" }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return "" }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemFlag(flagId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemFlag")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remFlag",
		[]string{"Log", "FlagId"},
		[]string{cmdContext.TestDemarcation(), flagId})
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemDockerImage(imageId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemDockerImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remDockerImage",
		[]string{"Log", "ImageId"},
		[]string{cmdContext.TestDemarcation(), imageId})
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallRemImageVersion(imageVersionId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallRemImageVersion")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"remImageVersion",
		[]string{"Log", "ImageVersionId"},
		[]string{cmdContext.TestDemarcation(), imageVersionId})
	if ! cmdContext.AssertErrIsNil(err, "") { return false}
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "") { return false }
	return responseMap, nil
}

/*******************************************************************************
 * Return an array of maps, each containing the fields on a DockerImageVersionDesc.
 */
func (cmdContext *CmdContext) CallGetDockerImageVersions(imageId string) ([]map[string]interface{}, error) {
	
	restContext.StartCall("CallGetDockerImageVersions")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"getDockerImageVersions",
		[]string{"Log", "DockerImageId"},
		[]string{cmdContext.TestDemarcation(), imageId})
	if ! cmdContext.AssertErrIsNil(err, "") { return nil}
	
	if ! cmdContext.Verify200Response(resp) {
		cmdContext.FailTest()
		return nil
	}

	var responseMaps []map[string]interface{}
	responseMaps, err = rest.ParseResponseBodyToPayloadMaps(resp.Body)
	if err != nil { fmt.Println(err.Error()); return nil }

	return responseMaps, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUpdateUserInfo(userId, userName,
	email string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallUpdateUserInfo")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"updateUserInfo",
		[]string{"Log", "UserId", "UserName", "EmailAddress"},
		[]string{cmdContext.TestDemarcation(), userId, userName, email})
	if ! cmdContext.AssertErrIsNil(err, "when calling updateUserInfo") { return }
	
	var responseMap map[string]interface{}
	cmdContext.Verify200Response(resp)
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUserExists(userId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallUserExists")

	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"userExists",
		[]string{"Log", "UserId"},
		[]string{cmdContext.TestDemarcation(), userId})
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	var responseMap map[string]interface{}
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallUseScanConfigForImage(dockerImageId,
	scanConfigId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallUseScanConfigForImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"useScanConfigForImage",
		[]string{"Log", "DockerImageId", "ScanConfigId"},
		[]string{cmdContext.TestDemarcation(), dockerImageId, scanConfigId})
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallStopUsingScanConfigForImage(dockerImageId,
	scanConfigId string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallStopUsingScanConfigForImage")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"stopUsingScanConfigForImage",
		[]string{"Log", "DockerImageId", "ScanConfigId"},
		[]string{cmdContext.TestDemarcation(), dockerImageId, scanConfigId})
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallEnableEmailVerification(enabled bool) (map[string]interface{}, error) {
	
	restContext.StartCall("CallEnableEmailVerification")
	
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
		[]string{cmdContext.TestDemarcation(), flag})
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil }
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallValidateAccountVerificationToken(token string) (map[string]interface{}, error) {
	
	restContext.StartCall("CallValidateAccountVerificationToken")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionPost(cmdContext.SessionId,
		"validateAccountVerificationToken",
		[]string{"Log", "AccountVerificationToken"},
		[]string{cmdContext.TestDemarcation(), token})
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	var responseMap map[string]interface{}
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil }
	
	return responseMap, nil
}

/*******************************************************************************
 * 
 */
func (cmdContext *CmdContext) CallClearAll() (map[string]interface{}, error) {
	
	restContext.StartCall("CallClearAll")
	
	var resp *http.Response
	var err error
	resp, err = cmdContext.SendSessionGet("",
		"clearAll",
		[]string{"Log"},
		[]string{cmdContext.TestDemarcation()})
	if ! cmdContext.AssertErrIsNil(err, "") { return }
	
	if ! cmdContext.Verify200Response(resp) { cmdContext.FailTest() }
	var responseMap map[string]interface{}
	responseMap, err = rest.ParseResponseBodyToMap(resp.Body)
	if ! cmdContext.AssertErrIsNil(err, "while parsing response body to map") { return nil }
	return responseMap, nil
}
