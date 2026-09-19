require "test_helper"

class ShippingMethodControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @shippingMethod = shippingMethods(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create shippingMethod" do
    assert_difference("ShippingMethod.count") do
      post shippingMethods_url, params: { shippingMethod: { name:"test string for name", flatRate:"test value", estimatedDays:100, asActive:true, MethodType:ShippingMethod.MethodTypes[0] } }
    end

    assert_redirected_to shippingMethods_url
  end

 
  
  test "should destroy shippingMethod" do
    assert_difference("ShippingMethod.count", -1) do
      delete shippingMethod_url(@shippingMethod)
    end

    assert_redirected_to shippingMethods_url
  end
  
end


