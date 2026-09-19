require "test_helper"

class ProductVariantControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productVariant = productVariants(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productVariant" do
    assert_difference("ProductVariant.count") do
      post productVariants_url, params: { productVariant: { sku:"test value", barcode:"test string for barcode", title:"test string for title", weight:"test value", requiresShipping:true, WeightUnit:ProductVariant.WeightUnits[0] } }
    end

    assert_redirected_to productVariants_url
  end

 
  
  test "should destroy productVariant" do
    assert_difference("ProductVariant.count", -1) do
      delete productVariant_url(@productVariant)
    end

    assert_redirected_to productVariants_url
  end
  
end


