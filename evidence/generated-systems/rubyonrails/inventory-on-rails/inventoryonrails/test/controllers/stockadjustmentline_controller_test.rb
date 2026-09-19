require "test_helper"

class StockAdjustmentLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @stockAdjustmentLine = stockAdjustmentLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create stockAdjustmentLine" do
    assert_difference("StockAdjustmentLine.count") do
      post stockAdjustmentLines_url, params: { stockAdjustmentLine: { lineNumber:100, quantity:"test value", UnitOfMeasure:StockAdjustmentLine.UnitOfMeasures[0], StockStatus:StockAdjustmentLine.StockStatuss[0] } }
    end

    assert_redirected_to stockAdjustmentLines_url
  end

 
  
  test "should destroy stockAdjustmentLine" do
    assert_difference("StockAdjustmentLine.count", -1) do
      delete stockAdjustmentLine_url(@stockAdjustmentLine)
    end

    assert_redirected_to stockAdjustmentLines_url
  end
  
end


