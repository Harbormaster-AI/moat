require "test_helper"

class WishlistItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @wishlistItem = wishlistItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create wishlistItem" do
    assert_difference("WishlistItem.count") do
      post wishlistItems_url, params: { wishlistItem: { addedDate:1.week.ago } }
    end

    assert_redirected_to wishlistItems_url
  end

 
  
  test "should destroy wishlistItem" do
    assert_difference("WishlistItem.count", -1) do
      delete wishlistItem_url(@wishlistItem)
    end

    assert_redirected_to wishlistItems_url
  end
  
end


