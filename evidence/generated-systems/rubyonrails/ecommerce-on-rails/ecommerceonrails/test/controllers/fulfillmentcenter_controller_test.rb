require "test_helper"

class FulfillmentCenterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @fulfillmentCenter = fulfillmentCenters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create fulfillmentCenter" do
    assert_difference("FulfillmentCenter.count") do
      post fulfillmentCenters_url, params: { fulfillmentCenter: { name:"test string for name", centerCode:"test string for centerCode", address:"test value", timezone:"test string for timezone", asActive:true } }
    end

    assert_redirected_to fulfillmentCenters_url
  end

 
  
  test "should destroy fulfillmentCenter" do
    assert_difference("FulfillmentCenter.count", -1) do
      delete fulfillmentCenter_url(@fulfillmentCenter)
    end

    assert_redirected_to fulfillmentCenters_url
  end
  
end


