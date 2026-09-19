require "test_helper"

class EvaluationMetricControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @evaluationMetric = evaluationMetrics(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create evaluationMetric" do
    assert_difference("EvaluationMetric.count") do
      post evaluationMetrics_url, params: { evaluationMetric: { name:"test string for name", value:"test value" } }
    end

    assert_redirected_to evaluationMetrics_url
  end

 
  
  test "should destroy evaluationMetric" do
    assert_difference("EvaluationMetric.count", -1) do
      delete evaluationMetric_url(@evaluationMetric)
    end

    assert_redirected_to evaluationMetrics_url
  end
  
end


