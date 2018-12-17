Feature: Cancel
    Scenario: Check cancel doesn't save frames
        Given I successfully run `chrono start something +pandas`
        When I run `chrono cancel`
        Then the output should match:
        """
        ^Cancelled project something \[pandas\] at \d{2}:\d{2}$
        """
        And a file named "appdir/frames.msgpack" should not exist
