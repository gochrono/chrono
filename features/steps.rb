Given(/I wait (\d*) seconds/) do |seconds|
    sleep seconds.to_i
end

Given(/I start tracking time for "([^"]*)"/) do |project|
    startAt = DateTime.now.strftime("%Y-%m-%d") + " 11:00"
    step %(I run `chrono start #{project} -v --at "#{startAt}"`)
end

Given(/I start tracking time with tag "([^"]*)" for "([^"]*)"/) do |tag, project|
    startAt = DateTime.now.strftime("%Y-%m-%d") + " 11:00"
    step %(I run `chrono start #{project} +#{tag} --at "#{startAt}"`)
end

Given(/I stop tracking time/) do
    startAt = DateTime.now.strftime("%Y-%m-%d") + " 12:00"
    step %(I run `chrono stop --at "#{startAt}"`)
end

Then(/the output should contain the current log line/) do
    startTime = DateTime.now.strftime("%A %e %B %Y")
    step %(the output should match /^#{startTime}$/)
end

Then(/the output should contain a frame line for "([^"]*)"/) do |project|
    step %(the output should match:
    """
    \t\(ID: [0-9a-z]{6}\) 11:00 to 12:00    1h 00m 00s  #{project}
    """)
end
