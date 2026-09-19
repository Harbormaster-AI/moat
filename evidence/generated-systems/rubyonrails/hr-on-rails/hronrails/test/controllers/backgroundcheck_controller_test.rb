require "test_helper"

class BackgroundCheckControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @backgroundCheck = backgroundChecks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create backgroundCheck" do
    assert_difference("BackgroundCheck.count") do
      post backgroundChecks_url, params: { backgroundCheck: { checkNumber:"test string for checkNumber", provider:"test string for provider", completedDate:1.week.ago, Status:BackgroundCheck.Statuss[0] } }
    end

    assert_redirected_to backgroundChecks_url
  end

 
  
  test "should destroy backgroundCheck" do
    assert_difference("BackgroundCheck.count", -1) do
      delete backgroundCheck_url(@backgroundCheck)
    end

    assert_redirected_to backgroundChecks_url
  end
  
end


