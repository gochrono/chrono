Feature: Log
    Scenario: Check log command shows correct output with one simple project
        Given I start tracking time for "something"
        And I stop tracking time
        When I run `chrono log`
        Then the output should contain the current log line
        And the output should match:
        """
        \t\(ID: [0-9a-z]{7}\) 11:00 to 12:00    1h 00m 00s  something
        """

    Scenario: Check log command shows correct output with one project with tags
        Given I start tracking time with tag "pandas" for "something"
        And I stop tracking time
        When I run `chrono log`
        Then the output should contain the current log line
        And the output should match:
        """
        \t\(ID: [0-9a-z]{7}\) 11:00 to 12:00    1h 00m 00s  something\s*\[pandas\]
        """
