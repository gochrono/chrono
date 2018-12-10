Feature: Status
    @announce-stderr
    @announce-stdout
    Scenario: Check no project message shows
        When I run `chrono status`
        Then the output should contain:
        """
        No project started
        """
