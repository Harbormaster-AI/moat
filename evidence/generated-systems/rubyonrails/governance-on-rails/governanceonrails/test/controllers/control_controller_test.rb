require "test_helper"

class ControlControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @control = controls(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create control" do
    assert_difference("Control.count") do
      post controls_url, params: { control: { name:"test string for name", objective:"test string for objective", ownerDepartment:"test string for ownerDepartment", ControlType:Control.ControlTypes[0], Frequency:Control.Frequencys[0], Status:Control.Statuss[0] } }
    end

    assert_redirected_to controls_url
  end

 
  
  test "should destroy control" do
    assert_difference("Control.count", -1) do
      delete control_url(@control)
    end

    assert_redirected_to controls_url
  end
  
end


