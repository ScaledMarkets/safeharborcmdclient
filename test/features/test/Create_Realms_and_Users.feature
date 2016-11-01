# language: en

Feature: Create Realms and Users
	
	@done
	Scenario: Verify that we can create a realm without being logged in first.
		Given that I am not logged into SafeHarbor
		When I call CreateRealmAnon
		Then the CreateRealmAnon HTTP response code should be 200
		And we can log in as the admin user that we just created.

	Scenario: Verify that one can retrieve the users of a realm.
		Given that there are 2 non-admin users in a realm,
		And an admin user is authenticated,
		When I retrieve the users of the realm,
		Then 3 users are returned
		And only 1 of those users is an admin user.
	
	