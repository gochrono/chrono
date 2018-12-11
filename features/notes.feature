Feature: Notes
    Scenario: Check note is shown when added through start -n
        Given I successfully run `chrono start something -n "I am a note"`
        When I run `chrono notes show`
        Then the output should match:
        """
        ^\[0\]: I am a note$
        """
    Scenario: Check note is shown when added through notes add
        Given I successfully run `chrono start something`
        And I successfully run `chrono notes add "I am a note"`
        When I run `chrono notes show`
        Then the output should match:
        """
        ^\[0\]: I am a note$
        """
    Scenario: Check multiple notes is shown when added through notes add
        Given I successfully run `chrono start something`
        And I successfully run `chrono notes add "I am a note"`
        And I successfully run `chrono notes add "I am still a note"`
        When I run `chrono notes show`
        Then the output should match:
        """
        ^\[0\]: I am a note$
        ^\[1\]: I am still a note$
        """
    Scenario: Check note can be deleted with notes delete
        Given I successfully run `chrono start something`
        And I successfully run `chrono notes add "I am a note"`
        When I run `chrono notes delete 0`
        And I run `chrono notes show`
        Then the output should not match:
        """
        ^\[0\]: I am a note$
        """
