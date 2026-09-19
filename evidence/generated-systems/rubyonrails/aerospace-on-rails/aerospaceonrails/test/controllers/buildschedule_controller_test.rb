require "test_helper"

class BuildScheduleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @buildSchedule = buildSchedules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create buildSchedule" do
    assert_difference("BuildSchedule.count") do
      post buildSchedules_url, params: { buildSchedule: { scheduleNumber:"test string for scheduleNumber", Status:BuildSchedule.Statuss[0] } }
    end

    assert_redirected_to buildSchedules_url
  end

 
  
  test "should destroy buildSchedule" do
    assert_difference("BuildSchedule.count", -1) do
      delete buildSchedule_url(@buildSchedule)
    end

    assert_redirected_to buildSchedules_url
  end
  
end


