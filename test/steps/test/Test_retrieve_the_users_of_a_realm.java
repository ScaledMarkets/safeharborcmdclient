package test;

import cucumber.api.Format;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import cucumber.api.java.en.And;

import org.json.*;

import static test.Utils.*;
import static test.Methods.*;

public class Test_retrieve_the_users_of_a_realm extends TestBase {
	
	String realm4AdminUserId = "realm4admin";
	String realm4AdminPswd = "RealmPswd";
	String realm4AdminUserName = "realm 4 Admin Full Name";
	String user1Id = "joe";
	String user1Name = "Joseph Smith";
	String user1Email = "joe@somewhere.com";
	String user1Pswd = "mypassword";
	String user2Id = "sally";
	String user2Name = "Sally Franklin";
	String user2Email = "sally@else.com";
	String user2Pswd = "miapswd";
	String realmId;

	String[] responses;

	@Given("^that there are 2 non-admin users in a realm,$")
	public void there_are_2_non_admin_users_in_a_realm() throws Exception {
		
		responses = makeRequest("CreateRealmAnon", "realm4", "realm 4 Org",
			realm4AdminUserId, realm4AdminUserName, "realm4admin@gmail.com",
			realm4AdminPswd);
			// Returns a UserDesc
		
		JSONObject jSONObject = parseResponses(responses);
		Object obj = jSONObject.get("RealmId");
		assertThat(obj instanceof String, responses[0]);
		realmId = (String)obj;
		
		makeRequest("CreateUser", user1Id, user1Name, user1Email, user1Pswd, realmId);
		makeRequest("CreateUser", user2Id, user2Name, user2Email, user2Pswd, realmId);
	}
	
	@And("^an admin user is authenticated,$")
	public void an_admin_user_is_authenticated() throws Exception {
		
		responses = makeRequest("GetMyDesc");
		JSONObject jSONObject = parseResponses(responses);
		Object obj = jSONObject.get("HTTPStatusCode");
		assertThat(obj instanceof Integer, responses[0]);
		int statusCode = ((Integer)obj).intValue();
		assertThat(statusCode != 200);
		
		obj = jSONObject.get("CanModifyTheseRealms");
		assertThat(obj instanceof String[], responses[0]);
		String[] canModifyRealms = (String[])obj;
		boolean foundRealm = false;
		for (String rid : canModifyRealms) {
			if (rid.equals(realmId)) break;
		}
		assertThat(foundRealm, responses[0]);
	}
	
	@When("^I retrieve the users of the realm,$")
	public void i_retrieve_the_users_of_the_realm() throws Exception {
		//responses = makeRequest(....);
	}

	@Then("^(\\d+) users are returned$")
	public void this_many_users_are_returned(int numUsers) throws Exception {
		
	}
	
	@And("^only (\\d+) of those users is an admin user\\.$")
	public void abc(int numAdminUsers) throws Exception {
		
	}
}
