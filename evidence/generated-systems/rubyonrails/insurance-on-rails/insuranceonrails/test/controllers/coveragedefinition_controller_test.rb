require "test_helper"

class CoverageDefinitionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @coverageDefinition = coverageDefinitions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create coverageDefinition" do
    assert_difference("CoverageDefinition.count") do
      post coverageDefinitions_url, params: { coverageDefinition: { name:"test string for name", defaultLimit:"test value", defaultDeductible:"test value", asMandatory:true, CoverageType:CoverageDefinition.CoverageTypes[0] } }
    end

    assert_redirected_to coverageDefinitions_url
  end

 
  
  test "should destroy coverageDefinition" do
    assert_difference("CoverageDefinition.count", -1) do
      delete coverageDefinition_url(@coverageDefinition)
    end

    assert_redirected_to coverageDefinitions_url
  end
  
end


