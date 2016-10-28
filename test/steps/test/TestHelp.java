package test;

import cucumber.api.Format;
import cucumber.api.Scenario;
import cucumber.api.java.Before;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;

import org.json.*;

import static test.Utils.*;

public class TestHelp extends TestBase {
	
	@When("^I request help$")
	public void i_request_help() throws Throwable {
		process = Runtime.getRuntime().exec(
			"bin/safeharborcmdclient -help");
	}
	
	@Then("^help is printed$")
	public void help_is_printed() throws Throwable {
		String[] str = Utils.getResponse(process);
		System.out.println("Obtained response: " + str[0]);
		System.out.println();
	}
}
