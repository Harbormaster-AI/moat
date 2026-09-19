require "test_helper"

class LeaveRequestControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @leaveRequest = leaveRequests(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create leaveRequest" do
    assert_difference("LeaveRequest.count") do
      post leaveRequests_url, params: { leaveRequest: { requestNumber:"test string for requestNumber", startDate:1.week.ago, endDate:1.week.ago, reason:"test string for reason", hours:"test value", Status:LeaveRequest.Statuss[0] } }
    end

    assert_redirected_to leaveRequests_url
  end

 
  
  test "should destroy leaveRequest" do
    assert_difference("LeaveRequest.count", -1) do
      delete leaveRequest_url(@leaveRequest)
    end

    assert_redirected_to leaveRequests_url
  end
  
end


