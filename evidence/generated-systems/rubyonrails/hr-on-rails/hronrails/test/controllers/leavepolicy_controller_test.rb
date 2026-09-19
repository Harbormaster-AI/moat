require "test_helper"

class LeavePolicyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @leavePolicy = leavePolicys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create leavePolicy" do
    assert_difference("LeavePolicy.count") do
      post leavePolicys_url, params: { leavePolicy: { name:"test string for name", accrualRate:"test value", carryoverAllowed:true, maxBalance:"test value", LeaveCategory:LeavePolicy.LeaveCategorys[0], AccrualUnit:LeavePolicy.AccrualUnits[0] } }
    end

    assert_redirected_to leavePolicys_url
  end

 
  
  test "should destroy leavePolicy" do
    assert_difference("LeavePolicy.count", -1) do
      delete leavePolicy_url(@leavePolicy)
    end

    assert_redirected_to leavePolicys_url
  end
  
end


