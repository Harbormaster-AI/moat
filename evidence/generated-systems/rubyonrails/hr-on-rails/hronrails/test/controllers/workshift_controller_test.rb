require "test_helper"

class WorkShiftControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @workShift = workShifts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create workShift" do
    assert_difference("WorkShift.count") do
      post workShifts_url, params: { workShift: { startTime:1.week.ago, endTime:1.week.ago, breakMinutes:100, DayOfWeek:WorkShift.DayOfWeeks[0] } }
    end

    assert_redirected_to workShifts_url
  end

 
  
  test "should destroy workShift" do
    assert_difference("WorkShift.count", -1) do
      delete workShift_url(@workShift)
    end

    assert_redirected_to workShifts_url
  end
  
end


