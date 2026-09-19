require "test_helper"

class PayoutControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @payout = payouts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create payout" do
    assert_difference("Payout.count") do
      post payouts_url, params: { payout: { payoutReference:"test string for payoutReference", amount:"test value", currency:"test string for currency", scheduledDate:1.week.ago, paidDate:1.week.ago, Status:Payout.Statuss[0] } }
    end

    assert_redirected_to payouts_url
  end

 
  
  test "should destroy payout" do
    assert_difference("Payout.count", -1) do
      delete payout_url(@payout)
    end

    assert_redirected_to payouts_url
  end
  
end


