require "test_helper"

class CartItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @cartItem = cartItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create cartItem" do
    assert_difference("CartItem.count") do
      post cartItems_url, params: { cartItem: { quantity:100, unitPrice:"test value", totalPrice:"test value" } }
    end

    assert_redirected_to cartItems_url
  end

 
  
  test "should destroy cartItem" do
    assert_difference("CartItem.count", -1) do
      delete cartItem_url(@cartItem)
    end

    assert_redirected_to cartItems_url
  end
  
end


