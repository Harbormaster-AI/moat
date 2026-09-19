require "test_helper"

class GiftCardControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @giftCard = giftCards(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create giftCard" do
    assert_difference("GiftCard.count") do
      post giftCards_url, params: { giftCard: { code:"test string for code", balance:"test value", expirationDate:1.week.ago, Status:GiftCard.Statuss[0] } }
    end

    assert_redirected_to giftCards_url
  end

 
  
  test "should destroy giftCard" do
    assert_difference("GiftCard.count", -1) do
      delete giftCard_url(@giftCard)
    end

    assert_redirected_to giftCards_url
  end
  
end


