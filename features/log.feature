Feature: Log
    Scenario: Check log command shows correct output with one simple project
        Given I successfully run `chrono start something --at "2018-12-11 11:00"`
        And I successfully run `chrono stop --at "2018-12-11 12:00"`
        When I run `chrono log`
        Then the output should match:
        """
        Tuesday 11 December 2018
        \t\(ID: [0-9a-z]{6}\) 11:00 to 12:00    1h 00m 00s  something
        """

    Scenario: Check log command shows correct output with project with tags
        Given I successfully run `chrono start something +cats --at "2018-12-11 11:00"`
        And I successfully run `chrono stop --at "2018-12-11 12:00"`
        When I run `chrono log`
        Then the output should match:
        """
        Tuesday 11 December 2018
        \t\(ID: [0-9a-z]{6}\) 11:00 to 12:00    1h 00m 00s  something\s*\[cats\]
        """
