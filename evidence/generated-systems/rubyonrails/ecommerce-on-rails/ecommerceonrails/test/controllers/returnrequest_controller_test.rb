require "test_helper"

class ReturnRequestControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @returnRequest = returnRequests(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create returnRequest" do
    assert_difference("ReturnRequest.count") do
      post returnRequests_url, params: { returnRequest: { returnNumber:"test string for returnNumber", createdAt:1.week.ago, refundAmount:"test value", Status:ReturnRequest.Statuss[0] } }
    end

    assert_redirected_to returnRequests_url
  end

 
  
  test "should destroy returnRequest" do
    assert_difference("ReturnRequest.count", -1) do
      delete returnRequest_url(@returnRequest)
    end

    assert_redirected_to returnRequests_url
  end
  
end


