require "test_helper"

class CartControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @cart = carts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create cart" do
    assert_difference("Cart.count") do
      post carts_url, params: { cart: { cartNumber:"test string for cartNumber", createdAt:1.week.ago, currency:"test string for currency", shippingAddress:"test value", billingAddress:"test value", Status:Cart.Statuss[0] } }
    end

    assert_redirected_to carts_url
  end

 
  
  test "should destroy cart" do
    assert_difference("Cart.count", -1) do
      delete cart_url(@cart)
    end

    assert_redirected_to carts_url
  end
  
end


