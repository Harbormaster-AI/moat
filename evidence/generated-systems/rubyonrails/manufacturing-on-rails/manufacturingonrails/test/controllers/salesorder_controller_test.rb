require "test_helper"

class SalesOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @salesOrder = salesOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create salesOrder" do
    assert_difference("SalesOrder.count") do
      post salesOrders_url, params: { salesOrder: { orderNumber:"test string for orderNumber", orderDate:1.week.ago, totalAmount:"test value", Status:SalesOrder.Statuss[0] } }
    end

    assert_redirected_to salesOrders_url
  end

 
  
  test "should destroy salesOrder" do
    assert_difference("SalesOrder.count", -1) do
      delete salesOrder_url(@salesOrder)
    end

    assert_redirected_to salesOrders_url
  end
  
end


