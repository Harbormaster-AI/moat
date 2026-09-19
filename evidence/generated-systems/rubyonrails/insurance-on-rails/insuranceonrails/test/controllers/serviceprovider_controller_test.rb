require "test_helper"

class ServiceProviderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @serviceProvider = serviceProviders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create serviceProvider" do
    assert_difference("ServiceProvider.count") do
      post serviceProviders_url, params: { serviceProvider: { name:"test string for name", taxId:"test string for taxId", ProviderType:ServiceProvider.ProviderTypes[0], NetworkStatus:ServiceProvider.NetworkStatuss[0] } }
    end

    assert_redirected_to serviceProviders_url
  end

 
  
  test "should destroy serviceProvider" do
    assert_difference("ServiceProvider.count", -1) do
      delete serviceProvider_url(@serviceProvider)
    end

    assert_redirected_to serviceProviders_url
  end
  
end


