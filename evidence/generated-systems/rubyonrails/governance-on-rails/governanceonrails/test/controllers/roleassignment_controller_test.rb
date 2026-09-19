require "test_helper"

class RoleAssignmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @roleAssignment = roleAssignments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create roleAssignment" do
    assert_difference("RoleAssignment.count") do
      post roleAssignments_url, params: { roleAssignment: { effectiveFrom:1.week.ago, effectiveTo:1.week.ago } }
    end

    assert_redirected_to roleAssignments_url
  end

 
  
  test "should destroy roleAssignment" do
    assert_difference("RoleAssignment.count", -1) do
      delete roleAssignment_url(@roleAssignment)
    end

    assert_redirected_to roleAssignments_url
  end
  
end


