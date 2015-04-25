require "pry"

$:.unshift File.dirname(__FILE__)
Dir[File.expand_path(File.join(__FILE__, "../**/*"))].each { |f| require f }

class ThriftClient
  attr_accessor :port

  def initialize(options={})
    @port = options.fetch(:port, 8000)
    transport.open
  end

  def transport
    @transport ||= Thrift::FramedTransport.new(Thrift::Socket.new("127.0.0.1"))
  end

  def protocol
    @protocol ||= Thrift::BinaryProtocol.new(transport)
  end

  def client
    @client ||= UsersService::Client.new(protocol)
  end

  def find(*args)
    client.find(*args)
  end

  def create(*args)
    client.create(*args)
  end

  def close
    transport.close
  end
end

class UsersModel
  attr_accessor :client, :first_name, :last_name, :email, :id

  class << self
    attr_accessor :client

    def find(*args)
      UsersModel.new(from_thrift @client.find(*args))
    end

    def create(options={})
      UsersModel.new(from_thrift @client.create(new_thrift(options)))
    end

    def new_thrift(options={})
      thrift_user = User.new
      options.each do |k, v|
        thrift_user.send("#{k}=", v)
      end

      thrift_user
    end

    def from_thrift(user)
      h = {}
      user.struct_fields.values.each do |field|
        h[field[:name].to_sym] = user.send(field[:name])
      end

      h
    end
  end

  @client = ThriftClient.new

  def initialize(options={})
    @id         = options.fetch(:id)
    @first_name = options.fetch(:first_name)
    @last_name  = options.fetch(:last_name)
    @email      = options.fetch(:email)
  end
end
