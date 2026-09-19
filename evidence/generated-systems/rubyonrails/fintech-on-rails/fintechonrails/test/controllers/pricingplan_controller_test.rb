require "test_helper"

class PricingPlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @pricingPlan = pricingPlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create pricingPlan" do
    assert_difference("PricingPlan.count") do
      post pricingPlans_url, params: { pricingPlan: { name:"test string for name", planCode:"test string for planCode", baseCurrency:"test string for baseCurrency", Status:PricingPlan.Statuss[0] } }
    end

    assert_redirected_to pricingPlans_url
  end

 
  
  test "should destroy pricingPlan" do
    assert_difference("PricingPlan.count", -1) do
      delete pricingPlan_url(@pricingPlan)
    end

    assert_redirected_to pricingPlans_url
  end
  
end


