require "test_helper"

class InboundShipmentLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inboundShipmentLine = inboundShipmentLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inboundShipmentLine" do
    assert_difference("InboundShipmentLine.count") do
      post inboundShipmentLines_url, params: { inboundShipmentLine: { lineNumber:100, quantity:"test value", UnitOfMeasure:InboundShipmentLine.UnitOfMeasures[0], StockStatus:InboundShipmentLine.StockStatuss[0] } }
    end

    assert_redirected_to inboundShipmentLines_url
  end

 
  
  test "should destroy inboundShipmentLine" do
    assert_difference("InboundShipmentLine.count", -1) do
      delete inboundShipmentLine_url(@inboundShipmentLine)
    end

    assert_redirected_to inboundShipmentLines_url
  end
  
end


