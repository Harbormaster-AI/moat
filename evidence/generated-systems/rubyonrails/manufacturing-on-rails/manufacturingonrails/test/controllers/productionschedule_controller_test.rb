require "test_helper"

class ProductionScheduleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productionSchedule = productionSchedules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productionSchedule" do
    assert_difference("ProductionSchedule.count") do
      post productionSchedules_url, params: { productionSchedule: { scheduleNumber:"test string for scheduleNumber", horizonStart:1.week.ago, horizonEnd:1.week.ago, Status:ProductionSchedule.Statuss[0] } }
    end

    assert_redirected_to productionSchedules_url
  end

 
  
  test "should destroy productionSchedule" do
    assert_difference("ProductionSchedule.count", -1) do
      delete productionSchedule_url(@productionSchedule)
    end

    assert_redirected_to productionSchedules_url
  end
  
end


