# language: en
Feature: Ping
	Scenario: Ping
		When I send a Ping request
		Then the HTTP response code should be 200
