package test;

import cucumber.api.Format;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;

import org.json.*;

public class TestPing {
	
	String host = System.getenv("SAFEHARBOR_HOST");
	int port = Integer.parseInt(System.getenv("SFEHARBOR_PORT"));
	Process process;
	
	@When("^I send a Ping request$")
	public void i_send_a_Ping_request() throws Throwable {
		process = Runtime.getRuntime().exec(
			"safeharbor -h " + host + " -p " + port + " Ping");
	}
	
	@Then("^the HTTP response code should be (\\d+)$")
	public void the_HTTP_response_code_should_be(int expected) throws Throwable {
		
		JSONObject json = Utils.getResponse(process);
		Object obj = json.get("HTTPStatusCode");
		assert obj instanceof Integer;
		assert ((Integer)obj).intValue() == expected;
	}
}
