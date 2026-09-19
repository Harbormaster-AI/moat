require "test_helper"

class TaxRuleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @taxRule = taxRules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create taxRule" do
    assert_difference("TaxRule.count") do
      post taxRules_url, params: { taxRule: { name:"test string for name", country:"test string for country", region:"test string for region", rate:"test value", taxInclusive:true, TaxClass:TaxRule.TaxClasss[0] } }
    end

    assert_redirected_to taxRules_url
  end

 
  
  test "should destroy taxRule" do
    assert_difference("TaxRule.count", -1) do
      delete taxRule_url(@taxRule)
    end

    assert_redirected_to taxRules_url
  end
  
end


