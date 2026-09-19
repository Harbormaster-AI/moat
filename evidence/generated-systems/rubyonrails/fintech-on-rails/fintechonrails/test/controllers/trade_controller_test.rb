require "test_helper"

class TradeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @trade = trades(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create trade" do
    assert_difference("Trade.count") do
      post trades_url, params: { trade: { executedAt:1.week.ago, quantity:"test value", price:"test value", fees:"test value", settlementDate:1.week.ago } }
    end

    assert_redirected_to trades_url
  end

 
  
  test "should destroy trade" do
    assert_difference("Trade.count", -1) do
      delete trade_url(@trade)
    end

    assert_redirected_to trades_url
  end
  
end


