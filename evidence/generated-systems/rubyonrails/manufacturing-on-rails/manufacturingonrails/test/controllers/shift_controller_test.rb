require "test_helper"

class ShiftControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @shift = shifts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create shift" do
    assert_difference("Shift.count") do
      post shifts_url, params: { shift: { shiftName:"test string for shiftName", startTime:"test string for startTime", endTime:"test string for endTime", ShiftType:Shift.ShiftTypes[0] } }
    end

    assert_redirected_to shifts_url
  end

 
  
  test "should destroy shift" do
    assert_difference("Shift.count", -1) do
      delete shift_url(@shift)
    end

    assert_redirected_to shifts_url
  end
  
end


