# language: en

Feature: Create Realms and Users
	
	@done
	Scenario: Verify that we can create a realm without being logged in first.
		Given that I am not logged into SafeHarbor
		When I call CreateRealmAnon
		Then the CreateRealmAnon HTTP response code should be 200
		And we can log in as the admin user that we just created.

	Scenario: Verify that one can retrieve the users of a realm.
		Given that admin user id realm4admin of realm4 is authenticated,
		And there are two non-admin users in the realm.
		When I retrieve the users of the realm,
		Then 3 users are returned
		And only one of those users is an admin user.
	
	