require "test_helper"

class ControlTest_ControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @controlTest_ = controlTest_s(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create controlTest_" do
    assert_difference("ControlTest_.count") do
      post controlTest_s_url, params: { controlTest_: { name:"test string for name", testPeriodStart:1.week.ago, testPeriodEnd:1.week.ago, sampleSize:100, TestType:ControlTest_.TestTypes[0], Effectiveness:ControlTest_.Effectivenesss[0], Status:ControlTest_.Statuss[0] } }
    end

    assert_redirected_to controlTest_s_url
  end

 
  
  test "should destroy controlTest_" do
    assert_difference("ControlTest_.count", -1) do
      delete controlTest__url(@controlTest_)
    end

    assert_redirected_to controlTest_s_url
  end
  
end


