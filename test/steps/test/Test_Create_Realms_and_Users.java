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
	String joeUserId = "jdoe";
	String joePswd = "weakpswd";
	
	String realm4Id;
	String user4AdminRealms;

	@Given("^that I am not logged into SafeHarbor$")
	public void that_I_am_not_logged_into_SafeHarbor() throws Throwable {
		
	}
	
	@When("^I send a Ping request$")
	public void () throws Throwable {
		process = makeRequest(CreateRealmAnon, "realm4", "realm 4 Org",
			realm4AdminUserId, "realm 4 Admin Full Name", "realm4admin@gmail.com",
			realm4AdminPswd);
	}

	@Then("^the HTTP response code should be (\\d+)$")
	public void the_HTTP_response_code_should_be(int expected) throws Throwable {
		
		JSONObject jSONObject = getResponseAsJSON(process);
		
		Object obj = jSONObject.get("realm4Id");
		assertThat(obj instanceof String);
		String realm4Id = (String)obj;
		
		obj = jSONObject.get("user4AdminRealms");
		assertThat(obj instanceof String[]);
		String[] user4AdminRealms = (String[])obj;
		
		assertThat(user4AdminRealms.length() == 1,
			"Wrong number of admin realms: " + user4AdminRealms.length());
	}
}
