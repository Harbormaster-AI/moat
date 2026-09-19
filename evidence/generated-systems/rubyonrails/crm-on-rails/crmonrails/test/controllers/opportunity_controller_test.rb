require "test_helper"

class OpportunityControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @opportunity = opportunitys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create opportunity" do
    assert_difference("Opportunity.count") do
      post opportunitys_url, params: { opportunity: { name:"test string for name", amount:"test value", closeDate:1.week.ago, probability:"test value", description:"test string for description", Stage:Opportunity.Stages[0], Type:Opportunity.Types[0], ForecastCategory:Opportunity.ForecastCategorys[0] } }
    end

    assert_redirected_to opportunitys_url
  end

 
  
  test "should destroy opportunity" do
    assert_difference("Opportunity.count", -1) do
      delete opportunity_url(@opportunity)
    end

    assert_redirected_to opportunitys_url
  end
  
end


