require "test_helper"

class TransferOrderLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @transferOrderLine = transferOrderLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create transferOrderLine" do
    assert_difference("TransferOrderLine.count") do
      post transferOrderLines_url, params: { transferOrderLine: { lineNumber:100, quantity:"test value", UnitOfMeasure:TransferOrderLine.UnitOfMeasures[0], StockStatus:TransferOrderLine.StockStatuss[0] } }
    end

    assert_redirected_to transferOrderLines_url
  end

 
  
  test "should destroy transferOrderLine" do
    assert_difference("TransferOrderLine.count", -1) do
      delete transferOrderLine_url(@transferOrderLine)
    end

    assert_redirected_to transferOrderLines_url
  end
  
end


