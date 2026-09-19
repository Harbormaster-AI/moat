require "test_helper"

class ConditionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @condition = conditions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create condition" do
    assert_difference("Condition.count") do
      post conditions_url, params: { condition: { code:"test string for code", onsetDate:1.week.ago, abatementDate:1.week.ago, ClinicalStatus:Condition.ClinicalStatuss[0], VerificationStatus:Condition.VerificationStatuss[0] } }
    end

    assert_redirected_to conditions_url
  end

 
  
  test "should destroy condition" do
    assert_difference("Condition.count", -1) do
      delete condition_url(@condition)
    end

    assert_redirected_to conditions_url
  end
  
end


