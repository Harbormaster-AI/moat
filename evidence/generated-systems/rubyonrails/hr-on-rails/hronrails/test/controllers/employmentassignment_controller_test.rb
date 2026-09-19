require "test_helper"

class EmploymentAssignmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @employmentAssignment = employmentAssignments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create employmentAssignment" do
    assert_difference("EmploymentAssignment.count") do
      post employmentAssignments_url, params: { employmentAssignment: { startDate:1.week.ago, endDate:1.week.ago, primary:true, AssignmentType:EmploymentAssignment.AssignmentTypes[0], Status:EmploymentAssignment.Statuss[0] } }
    end

    assert_redirected_to employmentAssignments_url
  end

 
  
  test "should destroy employmentAssignment" do
    assert_difference("EmploymentAssignment.count", -1) do
      delete employmentAssignment_url(@employmentAssignment)
    end

    assert_redirected_to employmentAssignments_url
  end
  
end


