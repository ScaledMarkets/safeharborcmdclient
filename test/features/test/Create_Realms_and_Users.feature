# language: en

Feature: Create Realms and Users
	
	@done
	Scenario: Verify that we can create a realm without being logged in first.
		Given that I am not logged into SafeHarbor
		When I call CreateRealmAnon
		Then the CreateRealmAnon HTTP response code should be 200
