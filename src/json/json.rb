require "quartz"

client = Quartz::Client.new(file_path: './src/json/json.go')

params = {'first_name' => 'Mark', 'last_name' => 'Bates'}
puts client.structs.inspect
puts client[:my_poster].struct_methods.inspect
resp = client[:my_poster].call('Echo', 'Params' => params)
puts resp.inspect
