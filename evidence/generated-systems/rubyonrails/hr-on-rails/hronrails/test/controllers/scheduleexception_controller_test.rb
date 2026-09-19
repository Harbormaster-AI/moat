require "test_helper"

class ScheduleExceptionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @scheduleException = scheduleExceptions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create scheduleException" do
    assert_difference("ScheduleException.count") do
      post scheduleExceptions_url, params: { scheduleException: { date:1.week.ago, reason:"test string for reason", hours:"test value" } }
    end

    assert_redirected_to scheduleExceptions_url
  end

 
  
  test "should destroy scheduleException" do
    assert_difference("ScheduleException.count", -1) do
      delete scheduleException_url(@scheduleException)
    end

    assert_redirected_to scheduleExceptions_url
  end
  
end


