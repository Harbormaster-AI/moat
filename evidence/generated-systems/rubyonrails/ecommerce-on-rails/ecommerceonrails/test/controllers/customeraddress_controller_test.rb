require "test_helper"

class CustomerAddressControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @customerAddress = customerAddresss(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create customerAddress" do
    assert_difference("CustomerAddress.count") do
      post customerAddresss_url, params: { customerAddress: { label:"test string for label", address:"test value", asDefaultShipping:true, asDefaultBilling:true } }
    end

    assert_redirected_to customerAddresss_url
  end

 
  
  test "should destroy customerAddress" do
    assert_difference("CustomerAddress.count", -1) do
      delete customerAddress_url(@customerAddress)
    end

    assert_redirected_to customerAddresss_url
  end
  
end


