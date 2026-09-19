require "test_helper"

class RiskControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @risk = risks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create risk" do
    assert_difference("Risk.count") do
      post risks_url, params: { risk: { name:"test string for name", description:"test string for description", inherentRiskScore:100, residualRiskScore:100, Category:Risk.Categorys[0], Impact:Risk.Impacts[0], Likelihood:Risk.Likelihoods[0], Status:Risk.Statuss[0] } }
    end

    assert_redirected_to risks_url
  end

 
  
  test "should destroy risk" do
    assert_difference("Risk.count", -1) do
      delete risk_url(@risk)
    end

    assert_redirected_to risks_url
  end
  
end


