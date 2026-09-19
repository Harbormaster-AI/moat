require "test_helper"

class ShiftAssignmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @shiftAssignment = shiftAssignments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create shiftAssignment" do
    assert_difference("ShiftAssignment.count") do
      post shiftAssignments_url, params: { shiftAssignment: { assignmentDate:1.week.ago } }
    end

    assert_redirected_to shiftAssignments_url
  end

 
  
  test "should destroy shiftAssignment" do
    assert_difference("ShiftAssignment.count", -1) do
      delete shiftAssignment_url(@shiftAssignment)
    end

    assert_redirected_to shiftAssignments_url
  end
  
end


