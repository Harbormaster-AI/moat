require "test_helper"

class MerchantControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @merchant = merchants(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create merchant" do
    assert_difference("Merchant.count") do
      post merchants_url, params: { merchant: { name:"test string for name", legalName:"test string for legalName", website:"test string for website", defaultCurrency:"test string for defaultCurrency", defaultLocale:"test string for defaultLocale", supportEmail:"test string for supportEmail" } }
    end

    assert_redirected_to merchants_url
  end

 
  
  test "should destroy merchant" do
    assert_difference("Merchant.count", -1) do
      delete merchant_url(@merchant)
    end

    assert_redirected_to merchants_url
  end
  
end


