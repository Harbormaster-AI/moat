require "test_helper"

class QuoteLineItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @quoteLineItem = quoteLineItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create quoteLineItem" do
    assert_difference("QuoteLineItem.count") do
      post quoteLineItems_url, params: { quoteLineItem: { quantity:"test value", unitPrice:"test value", discountAmount:"test value", taxAmount:"test value", totalAmount:"test value" } }
    end

    assert_redirected_to quoteLineItems_url
  end

 
  
  test "should destroy quoteLineItem" do
    assert_difference("QuoteLineItem.count", -1) do
      delete quoteLineItem_url(@quoteLineItem)
    end

    assert_redirected_to quoteLineItems_url
  end
  
end


