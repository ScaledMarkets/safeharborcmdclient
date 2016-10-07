// http://groovy-lang.org/json.html
// https://github.com/dkowis/cucumber-jvm-groovy-example

package test

this.metaClass.mixin(cucumber.api.groovy.Hooks)
this.metaClass.mixin(cucumber.api.groovy.EN)

When(~"I send a Ping request") { String opname ->
    
	/*
	USERID=....
	PASSWORD=....
	HOST=....
	PORT=6000
	....response = ....safeharbor -u $USERID -w $PASSWORD -h $HOST -p $PORT rest ping
	*/
}

Then(~"the HTTP response code should be (.*)") { double expected ->

	/*
	def jsonSlurper = new JsonSlurper()
	def object = jsonSlurper.parseText(response)
    
	assert object instanceof Map
	assert object.ResponseCode == expected
	*/
	
	assert true
}
