package test;

import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.BufferedReader;
import java.util.stream.Stream;

import org.json.*;
// https://github.com/stleary/JSON-java
// http://search.maven.org/#search%7Cga%7C1%7Cg%3A%22org.json%22

public class Utils {
	public static JSONObject getResponse(Process process) throws Exception {
		assert process != null;
		assert process.exitValue() == 0;
		
		boolean ok = process.waitFor(5, java.util.concurrent.TimeUnit.SECONDS);
		assert ok;
		InputStream os = process.getInputStream();
		assert os != null;
		BufferedReader lineReader = new BufferedReader(new InputStreamReader(os));
		String response = "";
		Stream<String> stream = lineReader.lines();
		String[] responseAr = { "" };
		stream.forEachOrdered(line -> responseAr[0] += line);
		
		JSONObject json = new JSONObject(response);
		return json;
	}
}
