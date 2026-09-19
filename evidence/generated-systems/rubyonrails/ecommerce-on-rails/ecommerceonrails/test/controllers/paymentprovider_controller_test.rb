require "test_helper"

class PaymentProviderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @paymentProvider = paymentProviders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create paymentProvider" do
    assert_difference("PaymentProvider.count") do
      post paymentProviders_url, params: { paymentProvider: { name:"test string for name", enabled:true, merchantAccountId:"test string for merchantAccountId", ProviderType:PaymentProvider.ProviderTypes[0] } }
    end

    assert_redirected_to paymentProviders_url
  end

 
  
  test "should destroy paymentProvider" do
    assert_difference("PaymentProvider.count", -1) do
      delete paymentProvider_url(@paymentProvider)
    end

    assert_redirected_to paymentProviders_url
  end
  
end


