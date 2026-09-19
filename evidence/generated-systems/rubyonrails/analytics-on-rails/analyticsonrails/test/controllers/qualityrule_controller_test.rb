require "test_helper"

class QualityRuleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @qualityRule = qualityRules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create qualityRule" do
    assert_difference("QualityRule.count") do
      post qualityRules_url, params: { qualityRule: { name:"test string for name", threshold:"test value", targetField:"test string for targetField", Dimension:QualityRule.Dimensions[0], Operator:QualityRule.Operators[0] } }
    end

    assert_redirected_to qualityRules_url
  end

 
  
  test "should destroy qualityRule" do
    assert_difference("QualityRule.count", -1) do
      delete qualityRule_url(@qualityRule)
    end

    assert_redirected_to qualityRules_url
  end
  
end


