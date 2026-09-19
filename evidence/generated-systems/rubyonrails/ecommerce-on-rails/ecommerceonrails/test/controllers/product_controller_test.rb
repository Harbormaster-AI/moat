require "test_helper"

class ProductControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @product = products(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create product" do
    assert_difference("Product.count") do
      post products_url, params: { product: { name:"test string for name", slug:"test string for slug", asActive:true, ProductType:Product.ProductTypes[0], DefaultTaxClass:Product.DefaultTaxClasss[0] } }
    end

    assert_redirected_to products_url
  end

 
  
  test "should destroy product" do
    assert_difference("Product.count", -1) do
      delete product_url(@product)
    end

    assert_redirected_to products_url
  end
  
end


