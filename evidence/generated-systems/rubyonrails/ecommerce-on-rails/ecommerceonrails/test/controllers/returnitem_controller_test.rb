require "test_helper"

class ReturnItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @returnItem = returnItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create returnItem" do
    assert_difference("ReturnItem.count") do
      post returnItems_url, params: { returnItem: { quantity:100, Reason:ReturnItem.Reasons[0], Condition:ReturnItem.Conditions[0] } }
    end

    assert_redirected_to returnItems_url
  end

 
  
  test "should destroy returnItem" do
    assert_difference("ReturnItem.count", -1) do
      delete returnItem_url(@returnItem)
    end

    assert_redirected_to returnItems_url
  end
  
end


