package test;

import cucumber.api.Format;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;

import org.json.*;

import static test.Utils.*;

public class TestPing extends TestBase {
	
	@When("^I send a Ping request$")
	public void i_send_a_Ping_request() throws Throwable {
		process = Runtime.getRuntime().exec(
			"bin/safeharborcmdclient -h " + getSafeHarborHost() + " -p " + getSafeHarborPort() + " Ping");
	}
	
	@Then("^the HTTP response code should be (\\d+)$")
	public void the_HTTP_response_code_should_be(int expected) throws Throwable {
		
		String str = Utils.getResponse(process);
		System.out.println("Obtained response: " + str);
		System.out.println();
		
		try {
			JSONObject json = new JSONObject(str);
			Object obj = json.get("HTTPStatusCode");
			assertThat(obj instanceof Integer);
			assertThat(((Integer)obj).intValue() == expected);
		} catch (Exception ex) {
			throw new Exception(ex.getMessage() + "; response=" + str);
		}
	}
}
