Feature: Status
    Scenario: No project has been started
        When I run `chrono status`
        Then the output should contain:
        """
        No project started
        """
    Scenario: Project has been started
        Given I successfully run `chrono start something`
        And I wait 1 seconds
        When I run `chrono status`
        Then the output should match /Project something started \d*\s(minutes?|seconds?) ago./

    Scenario: Format flag is used
        Given I successfully run `chrono start something`
        When I run `chrono status --format "{{ .Project }}"`
        Then the output should match /^something$/
