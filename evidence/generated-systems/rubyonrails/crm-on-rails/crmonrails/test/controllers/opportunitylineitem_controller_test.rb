require "test_helper"

class OpportunityLineItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @opportunityLineItem = opportunityLineItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create opportunityLineItem" do
    assert_difference("OpportunityLineItem.count") do
      post opportunityLineItems_url, params: { opportunityLineItem: { quantity:"test value", unitPrice:"test value", discountPercent:"test value", totalPrice:"test value" } }
    end

    assert_redirected_to opportunityLineItems_url
  end

 
  
  test "should destroy opportunityLineItem" do
    assert_difference("OpportunityLineItem.count", -1) do
      delete opportunityLineItem_url(@opportunityLineItem)
    end

    assert_redirected_to opportunityLineItems_url
  end
  
end


