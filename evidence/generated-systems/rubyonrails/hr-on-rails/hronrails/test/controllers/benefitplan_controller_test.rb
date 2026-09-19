require "test_helper"

class BenefitPlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @benefitPlan = benefitPlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create benefitPlan" do
    assert_difference("BenefitPlan.count") do
      post benefitPlans_url, params: { benefitPlan: { name:"test string for name", providerName:"test string for providerName", employeeContributionRate:"test value", employerContributionRate:"test value", eligibilityRules:"test string for eligibilityRules", BenefitType:BenefitPlan.BenefitTypes[0] } }
    end

    assert_redirected_to benefitPlans_url
  end

 
  
  test "should destroy benefitPlan" do
    assert_difference("BenefitPlan.count", -1) do
      delete benefitPlan_url(@benefitPlan)
    end

    assert_redirected_to benefitPlans_url
  end
  
end


