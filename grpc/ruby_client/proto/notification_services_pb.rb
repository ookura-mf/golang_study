# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: proto/notification.proto for package 'sample'

require 'grpc'
require './proto/notification_pb'

module Sample
  module Notifier
    # The greeting service definition.
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'sample.Notifier'

      # Sends a greeting
      rpc :PereodicHello, PereodicHelloRequest, stream(PeriodicHelloReply)
    end

    Stub = Service.rpc_stub_class
  end
end
