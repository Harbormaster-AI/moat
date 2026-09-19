require "test_helper"

class WishlistControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @wishlist = wishlists(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create wishlist" do
    assert_difference("Wishlist.count") do
      post wishlists_url, params: { wishlist: { name:"test string for name", asPublic:true, createdAt:1.week.ago } }
    end

    assert_redirected_to wishlists_url
  end

 
  
  test "should destroy wishlist" do
    assert_difference("Wishlist.count", -1) do
      delete wishlist_url(@wishlist)
    end

    assert_redirected_to wishlists_url
  end
  
end


