require "test_helper"

class TimeSeriesControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @timeSeries = timeSeriess(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create timeSeries" do
    assert_difference("TimeSeries.count") do
      post timeSeriess_url, params: { timeSeries: { name:"test string for name", timezone:"test string for timezone", Granularity:TimeSeries.Granularitys[0] } }
    end

    assert_redirected_to timeSeriess_url
  end

 
  
  test "should destroy timeSeries" do
    assert_difference("TimeSeries.count", -1) do
      delete timeSeries_url(@timeSeries)
    end

    assert_redirected_to timeSeriess_url
  end
  
end


