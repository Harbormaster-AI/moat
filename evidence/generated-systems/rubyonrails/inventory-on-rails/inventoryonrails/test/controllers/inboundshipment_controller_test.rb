require "test_helper"

class InboundShipmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inboundShipment = inboundShipments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inboundShipment" do
    assert_difference("InboundShipment.count") do
      post inboundShipments_url, params: { inboundShipment: { shipmentNumber:"test string for shipmentNumber", expectedArrivalDate:1.week.ago, arrivalDate:1.week.ago, carrierName:"test string for carrierName", Status:InboundShipment.Statuss[0] } }
    end

    assert_redirected_to inboundShipments_url
  end

 
  
  test "should destroy inboundShipment" do
    assert_difference("InboundShipment.count", -1) do
      delete inboundShipment_url(@inboundShipment)
    end

    assert_redirected_to inboundShipments_url
  end
  
end


