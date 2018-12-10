Feature: Start
    Scenario: Check start project
        Given a directory named "appdir"
        When I run `chrono start something`
        Then the output should match /^Starting project something at \d{2}:\d{2}/
        And a file named "appdir/state.msgpack" should exist

    Scenario: Check start --at project
        When I run `chrono start something --at 15:33`
        Then the output should contain "Starting project something at 15:33"
        And a file named "appdir/state.msgpack" should exist

    Scenario: --note flag is passed in
        When I run `chrono start something -n "a simple note"`
        Then I successfully run `chrono notes show`
        And the output should contain "[0]: a simple note"
