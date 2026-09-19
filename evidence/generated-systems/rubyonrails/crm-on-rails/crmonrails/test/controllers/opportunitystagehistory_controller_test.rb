require "test_helper"

class OpportunityStageHistoryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @opportunityStageHistory = opportunityStageHistorys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create opportunityStageHistory" do
    assert_difference("OpportunityStageHistory.count") do
      post opportunityStageHistorys_url, params: { opportunityStageHistory: { changedAt:1.week.ago, comment:"test string for comment", FromStage:OpportunityStageHistory.FromStages[0], ToStage:OpportunityStageHistory.ToStages[0] } }
    end

    assert_redirected_to opportunityStageHistorys_url
  end

 
  
  test "should destroy opportunityStageHistory" do
    assert_difference("OpportunityStageHistory.count", -1) do
      delete opportunityStageHistory_url(@opportunityStageHistory)
    end

    assert_redirected_to opportunityStageHistorys_url
  end
  
end


