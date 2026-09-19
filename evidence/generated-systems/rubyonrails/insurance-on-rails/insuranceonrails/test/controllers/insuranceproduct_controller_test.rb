require "test_helper"

class InsuranceProductControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @insuranceProduct = insuranceProducts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create insuranceProduct" do
    assert_difference("InsuranceProduct.count") do
      post insuranceProducts_url, params: { insuranceProduct: { name:"test string for name", productCode:"test string for productCode", LineOfBusiness:InsuranceProduct.LineOfBusinesss[0] } }
    end

    assert_redirected_to insuranceProducts_url
  end

 
  
  test "should destroy insuranceProduct" do
    assert_difference("InsuranceProduct.count", -1) do
      delete insuranceProduct_url(@insuranceProduct)
    end

    assert_redirected_to insuranceProducts_url
  end
  
end


