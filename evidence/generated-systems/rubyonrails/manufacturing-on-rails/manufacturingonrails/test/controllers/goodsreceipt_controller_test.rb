require "test_helper"

class GoodsReceiptControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @goodsReceipt = goodsReceipts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create goodsReceipt" do
    assert_difference("GoodsReceipt.count") do
      post goodsReceipts_url, params: { goodsReceipt: { receiptNumber:"test string for receiptNumber", receiptDate:1.week.ago, Status:GoodsReceipt.Statuss[0] } }
    end

    assert_redirected_to goodsReceipts_url
  end

 
  
  test "should destroy goodsReceipt" do
    assert_difference("GoodsReceipt.count", -1) do
      delete goodsReceipt_url(@goodsReceipt)
    end

    assert_redirected_to goodsReceipts_url
  end
  
end


