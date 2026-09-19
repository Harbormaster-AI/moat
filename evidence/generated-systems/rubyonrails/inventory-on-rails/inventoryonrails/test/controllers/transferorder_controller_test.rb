require "test_helper"

class TransferOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @transferOrder = transferOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create transferOrder" do
    assert_difference("TransferOrder.count") do
      post transferOrders_url, params: { transferOrder: { orderNumber:"test string for orderNumber", requestedShipDate:1.week.ago, requestedReceiveDate:1.week.ago, shippedDate:1.week.ago, receivedDate:1.week.ago, Status:TransferOrder.Statuss[0] } }
    end

    assert_redirected_to transferOrders_url
  end

 
  
  test "should destroy transferOrder" do
    assert_difference("TransferOrder.count", -1) do
      delete transferOrder_url(@transferOrder)
    end

    assert_redirected_to transferOrders_url
  end
  
end


