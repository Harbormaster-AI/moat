require "test_helper"

class ProductPricingControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productPricing = productPricings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productPricing" do
    assert_difference("ProductPricing.count") do
      post productPricings_url, params: { productPricing: { listPrice:"test value", salePrice:"test value", validFrom:1.week.ago, validTo:1.week.ago } }
    end

    assert_redirected_to productPricings_url
  end

 
  
  test "should destroy productPricing" do
    assert_difference("ProductPricing.count", -1) do
      delete productPricing_url(@productPricing)
    end

    assert_redirected_to productPricings_url
  end
  
end


