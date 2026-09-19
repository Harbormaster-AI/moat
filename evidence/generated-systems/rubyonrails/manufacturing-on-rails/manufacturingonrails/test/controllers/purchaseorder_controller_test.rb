require "test_helper"

class PurchaseOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @purchaseOrder = purchaseOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create purchaseOrder" do
    assert_difference("PurchaseOrder.count") do
      post purchaseOrders_url, params: { purchaseOrder: { poNumber:"test string for poNumber", orderDate:1.week.ago, totalAmount:"test value", Status:PurchaseOrder.Statuss[0] } }
    end

    assert_redirected_to purchaseOrders_url
  end

 
  
  test "should destroy purchaseOrder" do
    assert_difference("PurchaseOrder.count", -1) do
      delete purchaseOrder_url(@purchaseOrder)
    end

    assert_redirected_to purchaseOrders_url
  end
  
end


