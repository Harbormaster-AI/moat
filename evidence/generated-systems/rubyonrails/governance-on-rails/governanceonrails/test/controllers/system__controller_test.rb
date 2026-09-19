require "test_helper"

class System_ControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @system_ = system_s(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create system_" do
    assert_difference("System_.count") do
      post system_s_url, params: { system_: { name:"test string for name", ownerDepartment:"test string for ownerDepartment", SystemType:System_.SystemTypes[0] } }
    end

    assert_redirected_to system_s_url
  end

 
  
  test "should destroy system_" do
    assert_difference("System_.count", -1) do
      delete system__url(@system_)
    end

    assert_redirected_to system_s_url
  end
  
end


