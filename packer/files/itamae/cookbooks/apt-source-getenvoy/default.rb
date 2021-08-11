remote_file '/etc/gpg.8115BA8E629CC074.key' do
  owner 'root'
  group 'root'
  mode  '0644'

  notifies :run, 'execute[apt-key add /etc/gpg.8115BA8E629CC074.key]', :immediately
end

execute 'apt-key add /etc/gpg.8115BA8E629CC074.key' do
  action :nothing
end

remote_file '/etc/apt/sources.list.d/getenvoy.list' do
  owner 'root'
  group 'root'
  mode  '0644'
  notifies :run, 'execute[apt-get update]'
end
