package test;

import cucumber.api.Format;
import cucumber.api.Scenario;
import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;

import org.json.*;

public class TestHelp extends TestBase {
	
	@Before
	public void before(Scenario scenario) {
		setScenario(scenario);
	}
	
	@When("^I request help$")
	public void i_request_help() throws Throwable {
		getScenario().write("Executing...");
		process = Runtime.getRuntime().exec(
			"bin/safeharborcmdclient -help");
	}
	
	@Then("^help is printed$")
	public void help_is_printed() throws Throwable {
		String response = Utils.getResponse(process);
	}
}
