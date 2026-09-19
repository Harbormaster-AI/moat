require "test_helper"

class MRPRunControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @mRPRun = mRPRuns(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create mRPRun" do
    assert_difference("MRPRun.count") do
      post mRPRuns_url, params: { mRPRun: { runNumber:"test string for runNumber", runDateTime:1.week.ago, planningHorizonDays:100, Status:MRPRun.Statuss[0] } }
    end

    assert_redirected_to mRPRuns_url
  end

 
  
  test "should destroy mRPRun" do
    assert_difference("MRPRun.count", -1) do
      delete mRPRun_url(@mRPRun)
    end

    assert_redirected_to mRPRuns_url
  end
  
end


