require "test_helper"

class PriceBookControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @priceBook = priceBooks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create priceBook" do
    assert_difference("PriceBook.count") do
      post priceBooks_url, params: { priceBook: { name:"test string for name", asActive:true, description:"test string for description" } }
    end

    assert_redirected_to priceBooks_url
  end

 
  
  test "should destroy priceBook" do
    assert_difference("PriceBook.count", -1) do
      delete priceBook_url(@priceBook)
    end

    assert_redirected_to priceBooks_url
  end
  
end


