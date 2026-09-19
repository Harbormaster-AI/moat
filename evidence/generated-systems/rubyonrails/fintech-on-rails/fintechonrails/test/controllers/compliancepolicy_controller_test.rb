require "test_helper"

class CompliancePolicyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @compliancePolicy = compliancePolicys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create compliancePolicy" do
    assert_difference("CompliancePolicy.count") do
      post compliancePolicys_url, params: { compliancePolicy: { name:"test string for name", policyCode:"test string for policyCode", description:"test string for description", Status:CompliancePolicy.Statuss[0] } }
    end

    assert_redirected_to compliancePolicys_url
  end

 
  
  test "should destroy compliancePolicy" do
    assert_difference("CompliancePolicy.count", -1) do
      delete compliancePolicy_url(@compliancePolicy)
    end

    assert_redirected_to compliancePolicys_url
  end
  
end


