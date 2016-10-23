package test;

import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.BufferedReader;
import java.util.stream.Stream;

import org.json.*;
// https://github.com/stleary/JSON-java
// http://search.maven.org/#search%7Cga%7C1%7Cg%3A%22org.json%22

public class Utils {
	
	public static String getSafeHarborHost() {
		String host = System.getProperty("SAFEHARBOR_HOST");
		if (host == null) {
			throw new RuntimeException("Variable SAFEHARBOR_HOST undefined");
		}
		return host;
	}
	
	public static int getSafeHarborPort() {
		String portStr = System.getProperty("SAFEHARBOR_PORT");
		if (portStr == null) {
			throw new RuntimeException("Variable SFEHARBOR_PORT undefined");
		}
		int port = Integer.parseInt(portStr);
		return port;
	}
	
	public static String getResponse(Process process) throws Exception {
		assert process != null;
		assert process.exitValue() == 0;
		
		//boolean ok = process.waitFor(5, java.util.concurrent.TimeUnit.SECONDS);
		//assert ok;
		//if (! ok) { throw new Exception("timed out"); }
		InputStream os = process.getInputStream();
		assert os != null;
		BufferedReader lineReader = new BufferedReader(new InputStreamReader(os));
		Stream<String> stream = lineReader.lines();
		String[] responseAr = { "" };
		stream.forEachOrdered(line -> responseAr[0] += line);
		return responseAr[0];
	}
	
	public static JSONObject getResponseAsJSON(Process process) throws Exception {
		String response = getResponse(process);
		JSONObject json = new JSONObject(response);
		System.out.println("...done.");
		return json;
	}
	
	public static void assertThat(boolean expr) throws Exception {
		assertThat(expr, null);
	}
	
	public static void assertThat(boolean expr, String msg) throws Exception {
		if (msg != null) msg = "; " + msg;
		if (! expr) throw new Exception("Assertion violation" + msg);
	}
}
