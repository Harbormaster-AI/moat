require "test_helper"

class PlannedOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @plannedOrder = plannedOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create plannedOrder" do
    assert_difference("PlannedOrder.count") do
      post plannedOrders_url, params: { plannedOrder: { plannedOrderNumber:"test string for plannedOrderNumber", quantity:"test value", dueDate:1.week.ago, OrderType:PlannedOrder.OrderTypes[0], Status:PlannedOrder.Statuss[0] } }
    end

    assert_redirected_to plannedOrders_url
  end

 
  
  test "should destroy plannedOrder" do
    assert_difference("PlannedOrder.count", -1) do
      delete plannedOrder_url(@plannedOrder)
    end

    assert_redirected_to plannedOrders_url
  end
  
end


