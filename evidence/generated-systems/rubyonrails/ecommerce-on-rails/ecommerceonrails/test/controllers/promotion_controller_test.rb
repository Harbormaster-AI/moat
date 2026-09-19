require "test_helper"

class PromotionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @promotion = promotions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create promotion" do
    assert_difference("Promotion.count") do
      post promotions_url, params: { promotion: { name:"test string for name", code:"test string for code", value:"test value", startDate:1.week.ago, endDate:1.week.ago, asStackable:true, maxRedemptions:100, PromotionType:Promotion.PromotionTypes[0], DiscountType:Promotion.DiscountTypes[0] } }
    end

    assert_redirected_to promotions_url
  end

 
  
  test "should destroy promotion" do
    assert_difference("Promotion.count", -1) do
      delete promotion_url(@promotion)
    end

    assert_redirected_to promotions_url
  end
  
end


