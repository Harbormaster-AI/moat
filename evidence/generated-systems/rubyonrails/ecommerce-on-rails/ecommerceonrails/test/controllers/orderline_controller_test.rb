require "test_helper"

class OrderLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @orderLine = orderLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create orderLine" do
    assert_difference("OrderLine.count") do
      post orderLines_url, params: { orderLine: { quantity:100, unitPrice:"test value", totalPrice:"test value", taxRate:"test value", LineStatus:OrderLine.LineStatuss[0] } }
    end

    assert_redirected_to orderLines_url
  end

 
  
  test "should destroy orderLine" do
    assert_difference("OrderLine.count", -1) do
      delete orderLine_url(@orderLine)
    end

    assert_redirected_to orderLines_url
  end
  
end


