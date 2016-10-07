package test;

import cucumber.api.Format;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;

public class TestPing {
	
	@When("^I send a Ping request$")
	public void i_send_a_Ping_request() throws Throwable {
		// Write code here that turns the phrase above into concrete actions
		throw new Exception();
	}
	
	@Then("^the HTTP response code should be (\\d+)$")
	public void the_HTTP_response_code_should_be(int arg1) throws Throwable {
		// Write code here that turns the phrase above into concrete actions
		throw new Exception();
	}
}
