Feature: Stop
    Scenario: Check stop project when there is no state file
        When I run `chrono stop -v`
        Then the output should match /^No project started$/

    Scenario: Check stop project when there is an empty state file
        Given I successfully run `chrono start something`
        And I successfully run `chrono stop`
        When I run `chrono stop`
        Then the output should match /^No project started$/
