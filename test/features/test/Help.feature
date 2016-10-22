# language: en

Feature: Help
	
	Scenario: Check
		When I request help
		Then help is printed
