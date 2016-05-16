require 'quartz'

client = Quartz::Client.new(file_path: './src/adder/main.go')

puts client.structs.inspect

puts client[:my_adder].struct_methods.inspect

puts client[:my_adder].call('Add', {'A' => 2, 'B' => 5}).inspect
