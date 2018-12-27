Feature: Start
    Scenario: Check start project
        When I run `chrono start something`
        Then the output should match /^Starting project something at \d{2}:\d{2}/
        And a file named "appdir/state.msgpack" should exist

    Scenario: Check start project with tag
        When I run `chrono start something +pandas`
        Then the output should match /^Starting project something \[pandas\] at \d{2}:\d{2}/
        And a file named "appdir/state.msgpack" should exist

    Scenario: Check start --at project
        When I run `chrono start something --at 15:33`
        Then the output should contain "Starting project something at 15:33"
        And a file named "appdir/state.msgpack" should exist

    Scenario: Check start --at project with invalid time format
        When I run `chrono start something --at invalid`
        Then the output should contain "Invalid time format"
        And a file named "appdir/state.msgpack" should not exist

    Scenario: --note flag is passed in
        When I run `chrono start something -n "a simple note"`
        Then I successfully run `chrono notes show`
        And the output should contain "[0]: a simple note"

    Scenario: Check start project fails gracefully when existing project is started
        Given I successfully run `chrono start something`
        When I run `chrono start else`
        Then the output should match /^Project something is already started./
