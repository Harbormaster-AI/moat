require "test_helper"

class SubscriptionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @subscription = subscriptions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create subscription" do
    assert_difference("Subscription.count") do
      post subscriptions_url, params: { subscription: { subscriptionNumber:"test string for subscriptionNumber", nextBillingDate:1.week.ago, startDate:1.week.ago, endDate:1.week.ago, Status:Subscription.Statuss[0], Interval:Subscription.Intervals[0] } }
    end

    assert_redirected_to subscriptions_url
  end

 
  
  test "should destroy subscription" do
    assert_difference("Subscription.count", -1) do
      delete subscription_url(@subscription)
    end

    assert_redirected_to subscriptions_url
  end
  
end


