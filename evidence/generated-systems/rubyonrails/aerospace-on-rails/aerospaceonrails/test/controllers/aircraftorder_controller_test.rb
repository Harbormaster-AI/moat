require "test_helper"

class AircraftOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftOrder = aircraftOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftOrder" do
    assert_difference("AircraftOrder.count") do
      post aircraftOrders_url, params: { aircraftOrder: { orderNumber:"test string for orderNumber", totalAmount:"test value", Status:AircraftOrder.Statuss[0] } }
    end

    assert_redirected_to aircraftOrders_url
  end

 
  
  test "should destroy aircraftOrder" do
    assert_difference("AircraftOrder.count", -1) do
      delete aircraftOrder_url(@aircraftOrder)
    end

    assert_redirected_to aircraftOrders_url
  end
  
end


