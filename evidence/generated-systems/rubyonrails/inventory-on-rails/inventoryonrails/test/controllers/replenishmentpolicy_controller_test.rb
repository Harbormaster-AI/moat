require "test_helper"

class ReplenishmentPolicyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @replenishmentPolicy = replenishmentPolicys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create replenishmentPolicy" do
    assert_difference("ReplenishmentPolicy.count") do
      post replenishmentPolicys_url, params: { replenishmentPolicy: { minLevel:"test value", maxLevel:"test value", reorderPoint:"test value", reorderQuantity:"test value", leadTimeDays:100, reviewPeriodDays:100, PolicyType:ReplenishmentPolicy.PolicyTypes[0] } }
    end

    assert_redirected_to replenishmentPolicys_url
  end

 
  
  test "should destroy replenishmentPolicy" do
    assert_difference("ReplenishmentPolicy.count", -1) do
      delete replenishmentPolicy_url(@replenishmentPolicy)
    end

    assert_redirected_to replenishmentPolicys_url
  end
  
end


