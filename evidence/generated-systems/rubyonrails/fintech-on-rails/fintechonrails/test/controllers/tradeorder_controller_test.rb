require "test_helper"

class TradeOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @tradeOrder = tradeOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create tradeOrder" do
    assert_difference("TradeOrder.count") do
      post tradeOrders_url, params: { tradeOrder: { orderId:"test string for orderId", quantity:"test value", limitPrice:"test value", placedAt:1.week.ago, Side:TradeOrder.Sides[0], Type:TradeOrder.Types[0], Status:TradeOrder.Statuss[0], TimeInForce:TradeOrder.TimeInForces[0] } }
    end

    assert_redirected_to tradeOrders_url
  end

 
  
  test "should destroy tradeOrder" do
    assert_difference("TradeOrder.count", -1) do
      delete tradeOrder_url(@tradeOrder)
    end

    assert_redirected_to tradeOrders_url
  end
  
end


