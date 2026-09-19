require "test_helper"

class PriceBookEntryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @priceBookEntry = priceBookEntrys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create priceBookEntry" do
    assert_difference("PriceBookEntry.count") do
      post priceBookEntrys_url, params: { priceBookEntry: { unitPrice:"test value", effectiveDate:1.week.ago, expirationDate:1.week.ago, asActive:true } }
    end

    assert_redirected_to priceBookEntrys_url
  end

 
  
  test "should destroy priceBookEntry" do
    assert_difference("PriceBookEntry.count", -1) do
      delete priceBookEntry_url(@priceBookEntry)
    end

    assert_redirected_to priceBookEntrys_url
  end
  
end


