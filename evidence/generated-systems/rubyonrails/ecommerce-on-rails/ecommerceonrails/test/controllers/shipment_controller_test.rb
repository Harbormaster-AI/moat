require "test_helper"

class ShipmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @shipment = shipments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create shipment" do
    assert_difference("Shipment.count") do
      post shipments_url, params: { shipment: { shipmentNumber:"test string for shipmentNumber", shippedDate:1.week.ago, deliveredDate:1.week.ago, trackingNumber:"test string for trackingNumber", shippingAddress:"test value", Status:Shipment.Statuss[0], Carrier:Shipment.Carriers[0] } }
    end

    assert_redirected_to shipments_url
  end

 
  
  test "should destroy shipment" do
    assert_difference("Shipment.count", -1) do
      delete shipment_url(@shipment)
    end

    assert_redirected_to shipments_url
  end
  
end


