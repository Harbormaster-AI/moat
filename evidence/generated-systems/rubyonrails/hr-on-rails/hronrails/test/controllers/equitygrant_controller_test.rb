require "test_helper"

class EquityGrantControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @equityGrant = equityGrants(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create equityGrant" do
    assert_difference("EquityGrant.count") do
      post equityGrants_url, params: { equityGrant: { grantId:"test string for grantId", grantedUnits:100, vestingStart:1.week.ago, GrantType:EquityGrant.GrantTypes[0] } }
    end

    assert_redirected_to equityGrants_url
  end

 
  
  test "should destroy equityGrant" do
    assert_difference("EquityGrant.count", -1) do
      delete equityGrant_url(@equityGrant)
    end

    assert_redirected_to equityGrants_url
  end
  
end


