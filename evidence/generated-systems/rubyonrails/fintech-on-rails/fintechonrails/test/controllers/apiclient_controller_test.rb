require "test_helper"

class APIClientControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aPIClient = aPIClients(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aPIClient" do
    assert_difference("APIClient.count") do
      post aPIClients_url, params: { aPIClient: { name:"test string for name", clientId:"test string for clientId", redirectUri:"test string for redirectUri", ClientType:APIClient.ClientTypes[0] } }
    end

    assert_redirected_to aPIClients_url
  end

 
  
  test "should destroy aPIClient" do
    assert_difference("APIClient.count", -1) do
      delete aPIClient_url(@aPIClient)
    end

    assert_redirected_to aPIClients_url
  end
  
end


