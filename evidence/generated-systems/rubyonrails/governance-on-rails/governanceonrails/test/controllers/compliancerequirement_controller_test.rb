require "test_helper"

class ComplianceRequirementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @complianceRequirement = complianceRequirements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create complianceRequirement" do
    assert_difference("ComplianceRequirement.count") do
      post complianceRequirements_url, params: { complianceRequirement: { name:"test string for name", source:"test string for source", citation:"test string for citation", Applicability:ComplianceRequirement.Applicabilitys[0], Status:ComplianceRequirement.Statuss[0] } }
    end

    assert_redirected_to complianceRequirements_url
  end

 
  
  test "should destroy complianceRequirement" do
    assert_difference("ComplianceRequirement.count", -1) do
      delete complianceRequirement_url(@complianceRequirement)
    end

    assert_redirected_to complianceRequirements_url
  end
  
end


