require "test_helper"

class SellerControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @seller = sellers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create seller" do
    assert_difference("Seller.count") do
      post sellers_url, params: { seller: { name:"test string for name", sellerCode:"test string for sellerCode", contactEmail:"test string for contactEmail", Status:Seller.Statuss[0] } }
    end

    assert_redirected_to sellers_url
  end

 
  
  test "should destroy seller" do
    assert_difference("Seller.count", -1) do
      delete seller_url(@seller)
    end

    assert_redirected_to sellers_url
  end
  
end


