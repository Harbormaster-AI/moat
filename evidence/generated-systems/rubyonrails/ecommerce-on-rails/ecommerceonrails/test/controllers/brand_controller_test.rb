require "test_helper"

class BrandControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @brand = brands(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create brand" do
    assert_difference("Brand.count") do
      post brands_url, params: { brand: { name:"test string for name", description:"test string for description", website:"test string for website" } }
    end

    assert_redirected_to brands_url
  end

 
  
  test "should destroy brand" do
    assert_difference("Brand.count", -1) do
      delete brand_url(@brand)
    end

    assert_redirected_to brands_url
  end
  
end


