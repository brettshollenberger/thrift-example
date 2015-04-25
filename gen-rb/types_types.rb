#
# Autogenerated by Thrift Compiler (0.9.2)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#

require 'thrift'

module ErrorCode
  CLIENT_ERROR = 0
  NOT_FOUND = 1
  SERVER_ERROR = 2
  NOT_IMPLEMENTED_ERROR = 3
  VALUE_MAP = {0 => "CLIENT_ERROR", 1 => "NOT_FOUND", 2 => "SERVER_ERROR", 3 => "NOT_IMPLEMENTED_ERROR"}
  VALID_VALUES = Set.new([CLIENT_ERROR, NOT_FOUND, SERVER_ERROR, NOT_IMPLEMENTED_ERROR]).freeze
end

class ServiceException < ::Thrift::Exception
  include ::Thrift::Struct, ::Thrift::Struct_Union
  CODE = 1
  MESSAGE = 2

  FIELDS = {
    CODE => {:type => ::Thrift::Types::I32, :name => 'code', :enum_class => ::ErrorCode},
    MESSAGE => {:type => ::Thrift::Types::STRING, :name => 'message'}
  }

  def struct_fields; FIELDS; end

  def validate
    raise ::Thrift::ProtocolException.new(::Thrift::ProtocolException::UNKNOWN, 'Required field code is unset!') unless @code
    raise ::Thrift::ProtocolException.new(::Thrift::ProtocolException::UNKNOWN, 'Required field message is unset!') unless @message
    unless @code.nil? || ::ErrorCode::VALID_VALUES.include?(@code)
      raise ::Thrift::ProtocolException.new(::Thrift::ProtocolException::UNKNOWN, 'Invalid value of field code!')
    end
  end

  ::Thrift::Struct.generate_accessors self
end
