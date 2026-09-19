require "test_helper"

class SecurityControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @security = securitys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create security" do
    assert_difference("Security.count") do
      post securitys_url, params: { security: { symbol:"test string for symbol", isin:"test string for isin", cusip:"test string for cusip", currency:"test string for currency", SecurityType:Security.SecurityTypes[0] } }
    end

    assert_redirected_to securitys_url
  end

 
  
  test "should destroy security" do
    assert_difference("Security.count", -1) do
      delete security_url(@security)
    end

    assert_redirected_to securitys_url
  end
  
end


