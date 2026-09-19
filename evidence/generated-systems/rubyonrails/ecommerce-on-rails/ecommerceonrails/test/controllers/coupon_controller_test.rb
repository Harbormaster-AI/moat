require "test_helper"

class CouponControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @coupon = coupons(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create coupon" do
    assert_difference("Coupon.count") do
      post coupons_url, params: { coupon: { code:"test string for code", usageLimit:100, perCustomerLimit:100, expirationDate:1.week.ago, Status:Coupon.Statuss[0] } }
    end

    assert_redirected_to coupons_url
  end

 
  
  test "should destroy coupon" do
    assert_difference("Coupon.count", -1) do
      delete coupon_url(@coupon)
    end

    assert_redirected_to coupons_url
  end
  
end


