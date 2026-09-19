require "test_helper"

class InsurancePlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @insurancePlan = insurancePlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create insurancePlan" do
    assert_difference("InsurancePlan.count") do
      post insurancePlans_url, params: { insurancePlan: { name:"test string for name", planCode:"test string for planCode", PlanType:InsurancePlan.PlanTypes[0] } }
    end

    assert_redirected_to insurancePlans_url
  end

 
  
  test "should destroy insurancePlan" do
    assert_difference("InsurancePlan.count", -1) do
      delete insurancePlan_url(@insurancePlan)
    end

    assert_redirected_to insurancePlans_url
  end
  
end


