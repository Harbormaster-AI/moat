require "test_helper"

class TrainingRunControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @trainingRun = trainingRuns(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create trainingRun" do
    assert_difference("TrainingRun.count") do
      post trainingRuns_url, params: { trainingRun: { runLabel:"test string for runLabel", startedAt:1.week.ago, completedAt:1.week.ago, Status:TrainingRun.Statuss[0] } }
    end

    assert_redirected_to trainingRuns_url
  end

 
  
  test "should destroy trainingRun" do
    assert_difference("TrainingRun.count", -1) do
      delete trainingRun_url(@trainingRun)
    end

    assert_redirected_to trainingRuns_url
  end
  
end


