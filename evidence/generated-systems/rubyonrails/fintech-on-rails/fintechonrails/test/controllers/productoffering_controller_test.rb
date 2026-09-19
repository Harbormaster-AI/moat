require "test_helper"

class ProductOfferingControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productOffering = productOfferings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productOffering" do
    assert_difference("ProductOffering.count") do
      post productOfferings_url, params: { productOffering: { name:"test string for name", productCode:"test string for productCode", Category:ProductOffering.Categorys[0] } }
    end

    assert_redirected_to productOfferings_url
  end

 
  
  test "should destroy productOffering" do
    assert_difference("ProductOffering.count", -1) do
      delete productOffering_url(@productOffering)
    end

    assert_redirected_to productOfferings_url
  end
  
end


