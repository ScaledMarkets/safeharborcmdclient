package test;

import cucumber.api.Format;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;

import org.json.*;

import static test.Utils.*;

public class Test_Create_Realms_and_Users extends TestBase {
	
	String realm4AdminUserId = "realm4admin";
	String realm4AdminPswd = "RealmPswd";
	String realm4AdminUserName = "realm 4 Admin Full Name";
	
	String realm4Id;
	String user4AdminRealms;

	@Given("^that I am not logged into SafeHarbor$")
	public void that_I_am_not_logged_into_SafeHarbor() throws Throwable {
		
	}
	
	@When("^I send a Ping request$")
	public void () throws Throwable {
		process = makeRequest(CreateRealmAnon, "realm4", "realm 4 Org",
			realm4AdminUserId, realm4AdminUserName, "realm4admin@gmail.com",
			realm4AdminPswd);
	}

	@Then("^the HTTP response code should be (\\d+)$")
	public void the_HTTP_response_code_should_be(int expected) throws Throwable {
		
		JSONObject jSONObject = getResponseAsJSON(process);
		
		// Returns UserDesc, which contains:
		// Id string
		// UserId string
		// UserName string
		// RealmId string

		Object obj = jSONObject.get("Id");
		assertThat(obj instanceof String);
		String userObjId = (String)obj;
		
		obj = jSONObject.get("UserId");
		assertThat(obj instanceof String);
		String userId = (String)obj;
		
		obj = jSONObject.get("UserName");
		assertThat(obj instanceof String);
		String userName = (String)obj;
		
		obj = jSONObject.get("RealmId");
		assertThat(obj instanceof String);
		String realmId = (String)obj;
		
		assertThat(userId.equals(realm4admin));
		assertThat(userName.equals(realm4AdminUserName));
	}
}
