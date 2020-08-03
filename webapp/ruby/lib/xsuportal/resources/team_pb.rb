# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/resources/team.proto

require 'google/protobuf'

require 'xsuportal/resources/contestant_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("xsuportal/resources/team.proto", :syntax => :proto3) do
    add_message "xsuportal.proto.resources.Team" do
      optional :id, :int64, 1
      optional :name, :string, 2
      optional :leader_id, :string, 3
      repeated :member_ids, :string, 4
      optional :final_participation, :bool, 5
      optional :hidden, :bool, 6
      optional :withdrawn, :bool, 7
      optional :disqualified, :bool, 9
      optional :detail, :message, 8, "xsuportal.proto.resources.Team.TeamDetail"
      optional :leader, :message, 16, "xsuportal.proto.resources.Contestant"
      repeated :members, :message, 17, "xsuportal.proto.resources.Contestant"
    end
    add_message "xsuportal.proto.resources.Team.TeamDetail" do
      optional :email_address, :string, 1
      optional :invite_token, :string, 16
    end
  end
end

module Xsuportal
  module Proto
    module Resources
      Team = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.resources.Team").msgclass
      Team::TeamDetail = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("xsuportal.proto.resources.Team.TeamDetail").msgclass
    end
  end
end
