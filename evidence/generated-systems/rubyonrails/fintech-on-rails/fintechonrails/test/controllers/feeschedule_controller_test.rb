require "test_helper"

class FeeScheduleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @feeSchedule = feeSchedules(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create feeSchedule" do
    assert_difference("FeeSchedule.count") do
      post feeSchedules_url, params: { feeSchedule: { name:"test string for name", amount:"test value", percentage:"test value", minimum:"test value", maximum:"test value", FeeType:FeeSchedule.FeeTypes[0], CalculationMethod:FeeSchedule.CalculationMethods[0] } }
    end

    assert_redirected_to feeSchedules_url
  end

 
  
  test "should destroy feeSchedule" do
    assert_difference("FeeSchedule.count", -1) do
      delete feeSchedule_url(@feeSchedule)
    end

    assert_redirected_to feeSchedules_url
  end
  
end


