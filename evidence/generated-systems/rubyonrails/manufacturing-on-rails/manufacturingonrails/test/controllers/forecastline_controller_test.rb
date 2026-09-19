require "test_helper"

class ForecastLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @forecastLine = forecastLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create forecastLine" do
    assert_difference("ForecastLine.count") do
      post forecastLines_url, params: { forecastLine: { period:1.week.ago, quantity:"test value", confidence:"test value" } }
    end

    assert_redirected_to forecastLines_url
  end

 
  
  test "should destroy forecastLine" do
    assert_difference("ForecastLine.count", -1) do
      delete forecastLine_url(@forecastLine)
    end

    assert_redirected_to forecastLines_url
  end
  
end


