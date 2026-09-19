require "test_helper"

class ForecastControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @forecast = forecasts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create forecast" do
    assert_difference("Forecast.count") do
      post forecasts_url, params: { forecast: { forecastNumber:"test string for forecastNumber", forecastHorizonStart:1.week.ago, forecastHorizonEnd:1.week.ago, Method:Forecast.Methods[0] } }
    end

    assert_redirected_to forecasts_url
  end

 
  
  test "should destroy forecast" do
    assert_difference("Forecast.count", -1) do
      delete forecast_url(@forecast)
    end

    assert_redirected_to forecasts_url
  end
  
end


