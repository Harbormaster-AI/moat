require "test_helper"

class OrderItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @orderItem = orderItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create orderItem" do
    assert_difference("OrderItem.count") do
      post orderItems_url, params: { orderItem: { quantity:"test value", unitPrice:"test value", discountAmount:"test value", taxAmount:"test value", totalAmount:"test value" } }
    end

    assert_redirected_to orderItems_url
  end

 
  
  test "should destroy orderItem" do
    assert_difference("OrderItem.count", -1) do
      delete orderItem_url(@orderItem)
    end

    assert_redirected_to orderItems_url
  end
  
end


