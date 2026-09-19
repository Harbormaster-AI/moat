require "test_helper"

class OnboardingTaskControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @onboardingTask = onboardingTasks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create onboardingTask" do
    assert_difference("OnboardingTask.count") do
      post onboardingTasks_url, params: { onboardingTask: { taskNumber:"test string for taskNumber", name:"test string for name", dueDate:1.week.ago, Status:OnboardingTask.Statuss[0] } }
    end

    assert_redirected_to onboardingTasks_url
  end

 
  
  test "should destroy onboardingTask" do
    assert_difference("OnboardingTask.count", -1) do
      delete onboardingTask_url(@onboardingTask)
    end

    assert_redirected_to onboardingTasks_url
  end
  
end


