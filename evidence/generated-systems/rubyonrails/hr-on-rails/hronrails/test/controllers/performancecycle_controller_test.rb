require "test_helper"

class PerformanceCycleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @performanceCycle = performanceCycles(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create performanceCycle" do
    assert_difference("PerformanceCycle.count") do
      post performanceCycles_url, params: { performanceCycle: { name:"test string for name", startDate:1.week.ago, endDate:1.week.ago, Status:PerformanceCycle.Statuss[0] } }
    end

    assert_redirected_to performanceCycles_url
  end

 
  
  test "should destroy performanceCycle" do
    assert_difference("PerformanceCycle.count", -1) do
      delete performanceCycle_url(@performanceCycle)
    end

    assert_redirected_to performanceCycles_url
  end
  
end


