require "quartz"
require "benchmark"
require 'uri'
require 'net/http'


params = {'first_name' => 'Mark', 'last_name' => 'Bates'}
client = Quartz::Client.new(file_path: './src/concurrent/concurrent.go')

puts Benchmark.realtime {
  resp = client[:my_poster].call('MultiEcho', 'Params' => params)
}

uri = URI("http://quiet-waters-1228.herokuapp.com/echo.json")
threads = []

puts Benchmark.realtime {
  5.times do
    threads << Thread.new do
      res = Net::HTTP.post_form(uri, params)
    end
  end
  threads.map(&:join)
}
