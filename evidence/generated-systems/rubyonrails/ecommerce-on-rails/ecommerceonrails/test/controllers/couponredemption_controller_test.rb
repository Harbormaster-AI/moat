require "test_helper"

class CouponRedemptionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @couponRedemption = couponRedemptions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create couponRedemption" do
    assert_difference("CouponRedemption.count") do
      post couponRedemptions_url, params: { couponRedemption: { redeemedAt:1.week.ago } }
    end

    assert_redirected_to couponRedemptions_url
  end

 
  
  test "should destroy couponRedemption" do
    assert_difference("CouponRedemption.count", -1) do
      delete couponRedemption_url(@couponRedemption)
    end

    assert_redirected_to couponRedemptions_url
  end
  
end


