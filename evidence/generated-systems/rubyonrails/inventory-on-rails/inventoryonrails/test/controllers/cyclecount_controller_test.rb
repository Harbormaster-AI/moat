require "test_helper"

class CycleCountControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @cycleCount = cycleCounts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create cycleCount" do
    assert_difference("CycleCount.count") do
      post cycleCounts_url, params: { cycleCount: { countNumber:"test string for countNumber", scheduledDate:1.week.ago, performedDate:1.week.ago, approvedBy:"test string for approvedBy", Status:CycleCount.Statuss[0] } }
    end

    assert_redirected_to cycleCounts_url
  end

 
  
  test "should destroy cycleCount" do
    assert_difference("CycleCount.count", -1) do
      delete cycleCount_url(@cycleCount)
    end

    assert_redirected_to cycleCounts_url
  end
  
end


