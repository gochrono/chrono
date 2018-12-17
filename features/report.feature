Feature: Report
    Scenario: Check report command reports correct time for simple project with no tags
        Given I have project "something" that started at "10:00" and ended at "12:00"
        When I run `chrono report`
        Then there should be a current date timespan
        And the output should match:
        """
        ^something - 2h 00m 00s$
        """

    Scenario: Check report command reports correct time for project with tags
        Given I have project "something" with tags "+dev" that started at "10:00" and ended at "12:00"
        When I run `chrono report`
        Then there should be a current date timespan
        And the output should match:
        """
        ^something - 2h 00m 00s$
        ^\s*\[dev 2h 00m 00s\]$
        """

    Scenario: Check report command reports correct time for project with tags
        Given I have project "something" with tags "+dev +time" that started at "10:00" and ended at "12:00"
        When I run `chrono report`
        Then there should be a current date timespan
        And the output should match:
        """
        ^something - 2h 00m 00s$
        ^\s*\[dev 2h 00m 00s\]$
        ^\s*\[time 2h 00m 00s\]$
        """

    Scenario: Check report has correct totals for two similar projects with different tags
        Given I have project "something" with tags "+dev" that started at "10:00" and ended at "12:00"
        And I have project "something" with tags "+time" that started at "13:00" and ended at "14:00"
        When I run `chrono report`
        Then there should be a current date timespan
        And the output should match:
        """
        ^something - 3h 00m 00s$
        ^\s*\[dev 2h 00m 00s\]$
        ^\s*\[time 1h 00m 00s\]$
        """

    Scenario: Check report has correct totals for two different projects with same tags
        Given I have project "something" with tags "+dev" that started at "10:00" and ended at "12:00"
        And I have project "timevault" with tags "+dev" that started at "13:00" and ended at "14:00"
        When I run `chrono report`
        Then there should be a current date timespan
        And the output should match:
        """
        ^something - 2h 00m 00s$
        ^\s*\[dev 2h 00m 00s\]$
        """
        And the output should match:
        """
        ^timevault - 1h 00m 00s$
        ^\s*\[dev 1h 00m 00s\]$
        """
