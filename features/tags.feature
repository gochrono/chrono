Feature: Tags
    Scenario: Check single tag added by start command show up
        Given I successfully run `chrono start something +pandas`
        And I successfully run `chrono stop`
        When I run `chrono tags`
        Then the output should match /^pandas$/

    Scenario: Check multiple tags added by start command show up
        Given I successfully run `chrono start something +pandas +cats`
        And I successfully run `chrono stop`
        When I run `chrono tags`
        Then the output should match /^pandas$/
        And the output should match /^cats$/
