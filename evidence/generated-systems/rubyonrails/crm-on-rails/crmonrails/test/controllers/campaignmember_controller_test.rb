require "test_helper"

class CampaignMemberControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @campaignMember = campaignMembers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create campaignMember" do
    assert_difference("CampaignMember.count") do
      post campaignMembers_url, params: { campaignMember: { responded:true, Status:CampaignMember.Statuss[0], MemberType:CampaignMember.MemberTypes[0] } }
    end

    assert_redirected_to campaignMembers_url
  end

 
  
  test "should destroy campaignMember" do
    assert_difference("CampaignMember.count", -1) do
      delete campaignMember_url(@campaignMember)
    end

    assert_redirected_to campaignMembers_url
  end
  
end


