require "test_helper"

class GoodsReceiptLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @goodsReceiptLine = goodsReceiptLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create goodsReceiptLine" do
    assert_difference("GoodsReceiptLine.count") do
      post goodsReceiptLines_url, params: { goodsReceiptLine: { lineNumber:100, receivedQuantity:"test value", acceptedQuantity:"test value", rejectedQuantity:"test value", lot:"test value" } }
    end

    assert_redirected_to goodsReceiptLines_url
  end

 
  
  test "should destroy goodsReceiptLine" do
    assert_difference("GoodsReceiptLine.count", -1) do
      delete goodsReceiptLine_url(@goodsReceiptLine)
    end

    assert_redirected_to goodsReceiptLines_url
  end
  
end


