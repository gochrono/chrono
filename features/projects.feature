Feature: Projects
    Scenario: Check single project name added by start command show up
        Given I successfully run `chrono start something`
        And I successfully run `chrono stop`
        When I run `chrono projects`
        Then the output should match /^something$/

    Scenario: Check multiple projects added by two start commands show up
        Given I successfully run `chrono start something`
        And I successfully run `chrono stop`
        And I successfully run `chrono start domination`
        And I successfully run `chrono stop`
        When I run `chrono projects`
        Then the output should match /^domination$/
        And the output should match /^something$/
