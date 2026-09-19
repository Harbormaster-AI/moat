require "test_helper"

class UsageLimitControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @usageLimit = usageLimits(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create usageLimit" do
    assert_difference("UsageLimit.count") do
      post usageLimits_url, params: { usageLimit: { name:"test string for name", amount:"test value", count:100, Scope:UsageLimit.Scopes[0], Period:UsageLimit.Periods[0] } }
    end

    assert_redirected_to usageLimits_url
  end

 
  
  test "should destroy usageLimit" do
    assert_difference("UsageLimit.count", -1) do
      delete usageLimit_url(@usageLimit)
    end

    assert_redirected_to usageLimits_url
  end
  
end


