require "test_helper"

class WorkScheduleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @workSchedule = workSchedules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create workSchedule" do
    assert_difference("WorkSchedule.count") do
      post workSchedules_url, params: { workSchedule: { name:"test string for name", standardHoursPerWeek:"test value", ScheduleType:WorkSchedule.ScheduleTypes[0] } }
    end

    assert_redirected_to workSchedules_url
  end

 
  
  test "should destroy workSchedule" do
    assert_difference("WorkSchedule.count", -1) do
      delete workSchedule_url(@workSchedule)
    end

    assert_redirected_to workSchedules_url
  end
  
end


