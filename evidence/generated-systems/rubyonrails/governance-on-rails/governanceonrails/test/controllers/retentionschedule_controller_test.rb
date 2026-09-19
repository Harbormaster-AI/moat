require "test_helper"

class RetentionScheduleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @retentionSchedule = retentionSchedules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create retentionSchedule" do
    assert_difference("RetentionSchedule.count") do
      post retentionSchedules_url, params: { retentionSchedule: { name:"test string for name", retentionPeriodMonths:100, RetentionTrigger:RetentionSchedule.RetentionTriggers[0], DispositionAction:RetentionSchedule.DispositionActions[0], Status:RetentionSchedule.Statuss[0] } }
    end

    assert_redirected_to retentionSchedules_url
  end

 
  
  test "should destroy retentionSchedule" do
    assert_difference("RetentionSchedule.count", -1) do
      delete retentionSchedule_url(@retentionSchedule)
    end

    assert_redirected_to retentionSchedules_url
  end
  
end


