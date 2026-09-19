require "test_helper"

class AuthorizationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @authorization = authorizations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create authorization" do
    assert_difference("Authorization.count") do
      post authorizations_url, params: { authorization: { authNumber:"test string for authNumber", requestedService:"test string for requestedService", Status:Authorization.Statuss[0] } }
    end

    assert_redirected_to authorizations_url
  end

 
  
  test "should destroy authorization" do
    assert_difference("Authorization.count", -1) do
      delete authorization_url(@authorization)
    end

    assert_redirected_to authorizations_url
  end
  
end


