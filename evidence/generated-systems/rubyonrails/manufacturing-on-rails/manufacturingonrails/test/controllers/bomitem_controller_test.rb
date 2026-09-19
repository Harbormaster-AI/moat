require "test_helper"

class BOMItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @bOMItem = bOMItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create bOMItem" do
    assert_difference("BOMItem.count") do
      post bOMItems_url, params: { bOMItem: { lineNumber:100, quantity:"test value", scrapPercent:"test value" } }
    end

    assert_redirected_to bOMItems_url
  end

 
  
  test "should destroy bOMItem" do
    assert_difference("BOMItem.count", -1) do
      delete bOMItem_url(@bOMItem)
    end

    assert_redirected_to bOMItems_url
  end
  
end


