require "test_helper"

class DataSubjectRequestControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataSubjectRequest = dataSubjectRequests(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataSubjectRequest" do
    assert_difference("DataSubjectRequest.count") do
      post dataSubjectRequests_url, params: { dataSubjectRequest: { receivedDate:1.week.ago, dueDate:1.week.ago, requesterCountry:"test string for requesterCountry", RequestType:DataSubjectRequest.RequestTypes[0], Status:DataSubjectRequest.Statuss[0] } }
    end

    assert_redirected_to dataSubjectRequests_url
  end

 
  
  test "should destroy dataSubjectRequest" do
    assert_difference("DataSubjectRequest.count", -1) do
      delete dataSubjectRequest_url(@dataSubjectRequest)
    end

    assert_redirected_to dataSubjectRequests_url
  end
  
end


