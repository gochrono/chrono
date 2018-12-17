Given(/I wait (\d*) seconds/) do |seconds|
    sleep seconds.to_i
end

Given(/I have project "([^"]*)" that started at "([^"]*)" and ended at "([^"]*)"/) do |project, started, ended|
    currentDate = DateTime.now.strftime("%Y-%m-%d")
    startedAt = currentDate + " " + started
    endedAt = currentDate + " " + ended
    step %(I run `chrono start #{project} --at "#{startedAt}" --ended "#{endedAt}"`)
end

Given(/I have project "([^"]*)" with tags "([^"]*)" that started at "([^"]*)" and ended at "([^"]*)"/) do |project, tags, started, ended|
    currentDate = DateTime.now.strftime("%Y-%m-%d")
    startedAt = currentDate + " " + started
    endedAt = currentDate + " " + ended
    step %(I run `chrono start #{project} #{tags} --at "#{startedAt}" --ended "#{endedAt}"`)
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

Then(/there should be a current date timespan/) do
    startTime = DateTime.now.strftime("%a %e %B %Y")
    step %(the output should match /^#{startTime} -> #{startTime}$/)
end

Then(/the output should contain a frame line for "([^"]*)"/) do |project|
    step %(the output should match:
    """
    \t\(ID: [0-9a-z]{6}\) 11:00 to 12:00    1h 00m 00s  #{project}
    """)
end
