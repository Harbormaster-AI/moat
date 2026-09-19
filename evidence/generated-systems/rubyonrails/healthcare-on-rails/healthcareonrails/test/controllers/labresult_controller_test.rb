require "test_helper"

class LabResultControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @labResult = labResults(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create labResult" do
    assert_difference("LabResult.count") do
      post labResults_url, params: { labResult: { resultCode:"test string for resultCode", issuedDate:1.week.ago, Status:LabResult.Statuss[0] } }
    end

    assert_redirected_to labResults_url
  end

 
  
  test "should destroy labResult" do
    assert_difference("LabResult.count", -1) do
      delete labResult_url(@labResult)
    end

    assert_redirected_to labResults_url
  end
  
end


