require "test_helper"

class DataProcessingActivityControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataProcessingActivity = dataProcessingActivitys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataProcessingActivity" do
    assert_difference("DataProcessingActivity.count") do
      post dataProcessingActivitys_url, params: { dataProcessingActivity: { name:"test string for name", purpose:"test string for purpose", startDate:1.week.ago, LawfulBasis:DataProcessingActivity.LawfulBasiss[0] } }
    end

    assert_redirected_to dataProcessingActivitys_url
  end

 
  
  test "should destroy dataProcessingActivity" do
    assert_difference("DataProcessingActivity.count", -1) do
      delete dataProcessingActivity_url(@dataProcessingActivity)
    end

    assert_redirected_to dataProcessingActivitys_url
  end
  
end


