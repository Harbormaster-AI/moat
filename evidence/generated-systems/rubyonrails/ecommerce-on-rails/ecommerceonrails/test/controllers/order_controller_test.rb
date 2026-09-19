require "test_helper"

class OrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @order = orders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create order" do
    assert_difference("Order.count") do
      post orders_url, params: { order: { orderNumber:"test string for orderNumber", placedDate:1.week.ago, subtotal:"test value", discountTotal:"test value", shippingTotal:"test value", taxTotal:"test value", grandTotal:"test value", shippingAddress:"test value", billingAddress:"test value", Status:Order.Statuss[0] } }
    end

    assert_redirected_to orders_url
  end

 
  
  test "should destroy order" do
    assert_difference("Order.count", -1) do
      delete order_url(@order)
    end

    assert_redirected_to orders_url
  end
  
end


