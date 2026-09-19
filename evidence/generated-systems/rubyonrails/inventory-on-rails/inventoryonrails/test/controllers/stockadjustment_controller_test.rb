require "test_helper"

class StockAdjustmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @stockAdjustment = stockAdjustments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create stockAdjustment" do
    assert_difference("StockAdjustment.count") do
      post stockAdjustments_url, params: { stockAdjustment: { adjustmentNumber:"test string for adjustmentNumber", reason:"test string for reason", adjustmentDate:1.week.ago, AdjustmentType:StockAdjustment.AdjustmentTypes[0], Status:StockAdjustment.Statuss[0] } }
    end

    assert_redirected_to stockAdjustments_url
  end

 
  
  test "should destroy stockAdjustment" do
    assert_difference("StockAdjustment.count", -1) do
      delete stockAdjustment_url(@stockAdjustment)
    end

    assert_redirected_to stockAdjustments_url
  end
  
end


