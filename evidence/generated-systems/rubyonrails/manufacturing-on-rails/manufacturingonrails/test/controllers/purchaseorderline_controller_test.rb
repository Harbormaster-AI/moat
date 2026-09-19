require "test_helper"

class PurchaseOrderLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @purchaseOrderLine = purchaseOrderLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create purchaseOrderLine" do
    assert_difference("PurchaseOrderLine.count") do
      post purchaseOrderLines_url, params: { purchaseOrderLine: { lineNumber:100, quantity:"test value", unitPrice:"test value", dueDate:1.week.ago } }
    end

    assert_redirected_to purchaseOrderLines_url
  end

 
  
  test "should destroy purchaseOrderLine" do
    assert_difference("PurchaseOrderLine.count", -1) do
      delete purchaseOrderLine_url(@purchaseOrderLine)
    end

    assert_redirected_to purchaseOrderLines_url
  end
  
end


