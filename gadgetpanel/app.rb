require 'sinatra'

$stdout.sync = true
get '/' do
  $stdout.puts "Got it says STDOUT"
  puts "Got it sings puts"
  $stderr.puts "Got it screams STDERR"
  $stdout.puts "*****DB URL: #{ENV.fetch('DATABASE_URL')}"
  '{"pong": true}'
end
