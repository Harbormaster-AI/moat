require "test_helper"

class SalesOrderLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @salesOrderLine = salesOrderLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create salesOrderLine" do
    assert_difference("SalesOrderLine.count") do
      post salesOrderLines_url, params: { salesOrderLine: { lineNumber:100, quantity:"test value", unitPrice:"test value", dueDate:1.week.ago } }
    end

    assert_redirected_to salesOrderLines_url
  end

 
  
  test "should destroy salesOrderLine" do
    assert_difference("SalesOrderLine.count", -1) do
      delete salesOrderLine_url(@salesOrderLine)
    end

    assert_redirected_to salesOrderLines_url
  end
  
end


