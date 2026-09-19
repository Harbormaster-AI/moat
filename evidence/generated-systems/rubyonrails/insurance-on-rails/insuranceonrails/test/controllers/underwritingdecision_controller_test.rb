require "test_helper"

class UnderwritingDecisionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @underwritingDecision = underwritingDecisions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create underwritingDecision" do
    assert_difference("UnderwritingDecision.count") do
      post underwritingDecisions_url, params: { underwritingDecision: { notes:"test string for notes", decisionDate:1.week.ago, Decision:UnderwritingDecision.Decisions[0] } }
    end

    assert_redirected_to underwritingDecisions_url
  end

 
  
  test "should destroy underwritingDecision" do
    assert_difference("UnderwritingDecision.count", -1) do
      delete underwritingDecision_url(@underwritingDecision)
    end

    assert_redirected_to underwritingDecisions_url
  end
  
end


