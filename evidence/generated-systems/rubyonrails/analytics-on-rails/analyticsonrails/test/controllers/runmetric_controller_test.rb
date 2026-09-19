require "test_helper"

class RunMetricControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @runMetric = runMetrics(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create runMetric" do
    assert_difference("RunMetric.count") do
      post runMetrics_url, params: { runMetric: { name:"test string for name", value:"test value" } }
    end

    assert_redirected_to runMetrics_url
  end

 
  
  test "should destroy runMetric" do
    assert_difference("RunMetric.count", -1) do
      delete runMetric_url(@runMetric)
    end

    assert_redirected_to runMetrics_url
  end
  
end


