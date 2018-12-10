require 'aruba/cucumber'
require 'fileutils'
require 'tmpdir'

chrono_dir = Dir.mktmpdir('chrono_build')
raise 'chrono build failed' unless system("go build -o #{chrono_dir}/chrono")

Before do
    set_env 'CHRONO_CONFIG_DIR', File.expand_path(File.join(current_dir, "appdir"))
    set_env 'PATH', "#{chrono_dir}:#{ENV['PATH']}"
end
