require "test_helper"

class RefundControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @refund = refunds(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create refund" do
    assert_difference("Refund.count") do
      post refunds_url, params: { refund: { refundNumber:"test string for refundNumber", amount:"test value", reason:"test string for reason", createdAt:1.week.ago, Status:Refund.Statuss[0] } }
    end

    assert_redirected_to refunds_url
  end

 
  
  test "should destroy refund" do
    assert_difference("Refund.count", -1) do
      delete refund_url(@refund)
    end

    assert_redirected_to refunds_url
  end
  
end


