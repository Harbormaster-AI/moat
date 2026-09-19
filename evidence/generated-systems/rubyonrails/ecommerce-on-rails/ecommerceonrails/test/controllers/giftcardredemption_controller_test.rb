require "test_helper"

class GiftCardRedemptionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @giftCardRedemption = giftCardRedemptions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create giftCardRedemption" do
    assert_difference("GiftCardRedemption.count") do
      post giftCardRedemptions_url, params: { giftCardRedemption: { redeemedAt:1.week.ago, amount:"test value" } }
    end

    assert_redirected_to giftCardRedemptions_url
  end

 
  
  test "should destroy giftCardRedemption" do
    assert_difference("GiftCardRedemption.count", -1) do
      delete giftCardRedemption_url(@giftCardRedemption)
    end

    assert_redirected_to giftCardRedemptions_url
  end
  
end


