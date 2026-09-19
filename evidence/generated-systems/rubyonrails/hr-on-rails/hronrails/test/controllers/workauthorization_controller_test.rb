require "test_helper"

class WorkAuthorizationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @workAuthorization = workAuthorizations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create workAuthorization" do
    assert_difference("WorkAuthorization.count") do
      post workAuthorizations_url, params: { workAuthorization: { country:"test string for country", expirationDate:1.week.ago, Status:WorkAuthorization.Statuss[0] } }
    end

    assert_redirected_to workAuthorizations_url
  end

 
  
  test "should destroy workAuthorization" do
    assert_difference("WorkAuthorization.count", -1) do
      delete workAuthorization_url(@workAuthorization)
    end

    assert_redirected_to workAuthorizations_url
  end
  
end


