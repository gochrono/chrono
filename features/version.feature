Feature: Version
    Scenario: Check version output is correct
        When I run `chrono version`
        Then the output should contain:
        """
        Version: dev
        Commit: none
        Built: unknown
        """
