Feature: Restart
    Scenario: Check restart project
        Given I successfully run `chrono start something +pandas`
        And I successfully run `chrono stop`
        When I run `chrono restart`
        Then the output should match:
        """
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        ^Stopping project something \[pandas\], started (\d{1,2} seconds? ago|now) \(id: [a-z0-9]{7}\)$
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        """
        And a file named "appdir/state.msgpack" should exist
    Scenario: Check restart project with --at flag
        Given I successfully run `chrono start something +pandas`
        And I successfully run `chrono stop`
        When I run `chrono restart --at 12:00`
        Then the output should match:
        """
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        ^Stopping project something \[pandas\], started (\d{1,2} seconds? ago|now) \(id: [a-z0-9]{7}\)$
        ^Starting project something \[pandas\] at 12:00$
        """
        And a file named "appdir/state.msgpack" should exist
    Scenario: Check restart project fails if invalid time given to --at flag
        Given I successfully run `chrono start something +pandas`
        And I successfully run `chrono stop`
        When I run `chrono restart --at invalid`
        Then the output should match:
        """
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        ^Stopping project something \[pandas\], started (\d{1,2} seconds? ago|now) \(id: [a-z0-9]{7}\)$
        ^Invalid time format$
        """
        And a file named "appdir/state.msgpack" should exist
    Scenario: Check restart project with zero index target
        Given I successfully run `chrono start something +pandas`
        And I successfully run `chrono stop`
        When I run `chrono restart 0`
        Then the output should match:
        """
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        ^Stopping project something \[pandas\], started (\d{1,2} seconds? ago|now) \(id: [a-z0-9]{7}\)$
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        """
        And a file named "appdir/state.msgpack" should exist
    Scenario: Check restart project with -2 index target
        Given I successfully run `chrono start something +pandas`
        And I successfully run `chrono stop`
        And I successfully run `chrono start else +pandas`
        And I successfully run `chrono stop`
        When I run `chrono restart -- -2`
        Then the output should match:
        """
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        ^Stopping project something \[pandas\], started (\d{1,2} seconds? ago|now) \(id: [a-z0-9]{7}\)$
        ^Starting project else \[pandas\] at \d{2}:\d{2}$
        ^Stopping project else \[pandas\], started (\d{1,2} seconds? ago|now) \(id: [a-z0-9]{7}\)$
        ^Starting project something \[pandas\] at \d{2}:\d{2}$
        """
        And a file named "appdir/state.msgpack" should exist

    Scenario: Check restart project fails gracefully when existing project is started
        Given I successfully run `chrono start something`
        When I run `chrono restart`
        Then the output should match /^Project something is already started./
