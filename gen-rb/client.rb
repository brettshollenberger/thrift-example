require "pry"

$:.unshift File.dirname(__FILE__)
Dir[File.expand_path(File.join(__FILE__, "../**/*"))].each { |f| require f }

begin
	port = ARGV[0] || 8000

  transport = Thrift::FramedTransport.new(Thrift::Socket.new("127.0.0.1"))
	protocol = Thrift::BinaryProtocol.new(transport)
	client = UsersService::Client.new(protocol)

	transport.open()

	result = client.find()
	puts result.inspect
	
	transport.close()
rescue
	puts $!
end
