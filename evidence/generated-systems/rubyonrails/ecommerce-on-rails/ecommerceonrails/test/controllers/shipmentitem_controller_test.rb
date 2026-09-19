require "test_helper"

class ShipmentItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @shipmentItem = shipmentItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create shipmentItem" do
    assert_difference("ShipmentItem.count") do
      post shipmentItems_url, params: { shipmentItem: { quantity:100 } }
    end

    assert_redirected_to shipmentItems_url
  end

 
  
  test "should destroy shipmentItem" do
    assert_difference("ShipmentItem.count", -1) do
      delete shipmentItem_url(@shipmentItem)
    end

    assert_redirected_to shipmentItems_url
  end
  
end


