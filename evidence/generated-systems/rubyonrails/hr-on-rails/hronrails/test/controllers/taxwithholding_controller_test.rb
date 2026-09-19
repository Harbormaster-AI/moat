require "test_helper"

class TaxWithholdingControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @taxWithholding = taxWithholdings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create taxWithholding" do
    assert_difference("TaxWithholding.count") do
      post taxWithholdings_url, params: { taxWithholding: { taxId:"test value", allowances:100, additionalAmount:"test value", FilingStatus:TaxWithholding.FilingStatuss[0] } }
    end

    assert_redirected_to taxWithholdings_url
  end

 
  
  test "should destroy taxWithholding" do
    assert_difference("TaxWithholding.count", -1) do
      delete taxWithholding_url(@taxWithholding)
    end

    assert_redirected_to taxWithholdings_url
  end
  
end


