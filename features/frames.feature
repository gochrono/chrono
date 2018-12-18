Feature: Frames
    Scenario: Check frames command shows correct output with one simple project
        Given I successfully run `chrono start something`
        And I successfully run `chrono stop`
        When I run `chrono frames`
        Then the output should match:
        """
        ^[0-9a-z]{7}$
        """
