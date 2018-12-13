# vi: ft=ruby

Vagrant.configure("2") do |config|
    config.vm.define "arch" do |arch|
        arch.vm.box = "archlinux/archlinux"
        arch.vm.hostname = 'arch'
        arch.vm.provider :virtualbox do |v|
            v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        end
    end

    config.vm.define "ubuntu" do |ubuntu|
        ubuntu.vm.box = "ubuntu/xenial64"
        ubuntu.vm.hostname = 'arch'
        ubuntu.vm.provider :virtualbox do |v|
            v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        end
    end

end
