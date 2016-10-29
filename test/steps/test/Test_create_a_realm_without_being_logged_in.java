package test;

import cucumber.api.Format;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import cucumber.api.java.en.And;

import org.json.*;

import static test.Utils.*;

public class Test_create_a_realm_without_being_logged_in extends TestBase {
	
	String realm4AdminUserId = "realm4admin";
	String realm4AdminPswd = "RealmPswd";
	String realm4AdminUserName = "realm 4 Admin Full Name";
	
	String realm4Id;
	String user4AdminRealms;
	String[] responses;

	@Given("^that I am not logged into SafeHarbor$")
	public void that_I_am_not_logged_into_SafeHarbor() throws Exception {
		
	}
	
	@When("^I call CreateRealmAnon$")
	public void i_call_CreateRealmAnon() throws Exception {
		responses = makeRequest("CreateRealmAnon", "realm4", "realm 4 Org",
			realm4AdminUserId, realm4AdminUserName, "realm4admin@gmail.com",
			realm4AdminPswd);
	}

	@Then("^the CreateRealmAnon HTTP response code should be (\\d+)$")
	public void the_CreateRealmAnon_HTTP_response_code_should_be(int expected) throws Exception {
		
		// Returns UserDesc, which contains:
		// Id string
		// UserId string
		// UserName string
		// RealmId string

		JSONObject jSONObject;
		try {
			jSONObject = new JSONObject(responses[0]);
		} catch (Exception ex) {
			throw new Exception("stdout=" + responses[0] + ", stderr=" + responses[1], ex);
		}
		
		Object obj = jSONObject.get("Id");
		assertThat(obj instanceof String, responses[0]);
		String userObjId = (String)obj;
		
		obj = jSONObject.get("UserId");
		assertThat(obj instanceof String, responses[0]);
		String userId = (String)obj;
		
		obj = jSONObject.get("Name");
		assertThat(obj instanceof String, responses[0]);
		String userName = (String)obj;
		
		obj = jSONObject.get("RealmId");
		assertThat(obj instanceof String, responses[0]);
		String realmId = (String)obj;
		
		assertThat(userId.equals(realm4AdminUserId));
		assertThat(userName.equals(realm4AdminUserName));
	}
	
	// Verify that we can log in as the admin user that we just created.
	@And("^we can log in as the admin user that we just created\\.$")
	public void we_can_log_in_as_the_admin_user_that_we_just_created() throws Exception {
		
		responses = makeRequest("Authenticate", realm4AdminUserId, realm4AdminPswd);
		
		JSONObject jSONObject;
		try {
			jSONObject = new JSONObject(responses[0]);
		} catch (Exception ex) {
			throw new Exception("stdout=" + responses[0] + ", stderr=" + responses[1], ex);
		}
		
		Object obj = jSONObject.get("AuthenticatedUserid");
		assertThat(obj instanceof String, responses[0]);
		String retUserId = (String)obj;
		assertThat(retUserId.equals(realm4AdminUserId), responses[0]);

		obj = jSONObject.get("IsAdmin");
		assertThat(obj instanceof Boolean, responses[0]);
		boolean retIsAdmin = (Boolean)obj;
		assertThat(retIsAdmin, responses[0]);
	}
}
