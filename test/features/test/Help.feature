# language: en

Feature: Help
	
	@done
	Scenario: Check
		When I request help
		Then help is printed
