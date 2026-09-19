require "test_helper"

class FXQuoteControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @fXQuote = fXQuotes(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create fXQuote" do
    assert_difference("FXQuote.count") do
      post fXQuotes_url, params: { fXQuote: { baseCurrency:"test string for baseCurrency", quoteCurrency:"test string for quoteCurrency", rate:"test value", quotedAt:1.week.ago, expiresAt:1.week.ago, PriceType:FXQuote.PriceTypes[0] } }
    end

    assert_redirected_to fXQuotes_url
  end

 
  
  test "should destroy fXQuote" do
    assert_difference("FXQuote.count", -1) do
      delete fXQuote_url(@fXQuote)
    end

    assert_redirected_to fXQuotes_url
  end
  
end


