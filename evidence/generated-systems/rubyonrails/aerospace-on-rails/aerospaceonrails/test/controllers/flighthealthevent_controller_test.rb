require "test_helper"

class FlightHealthEventControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @flightHealthEvent = flightHealthEvents(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create flightHealthEvent" do
    assert_difference("FlightHealthEvent.count") do
      post flightHealthEvents_url, params: { flightHealthEvent: { eventCode:"test string for eventCode", Severity:FlightHealthEvent.Severitys[0] } }
    end

    assert_redirected_to flightHealthEvents_url
  end

 
  
  test "should destroy flightHealthEvent" do
    assert_difference("FlightHealthEvent.count", -1) do
      delete flightHealthEvent_url(@flightHealthEvent)
    end

    assert_redirected_to flightHealthEvents_url
  end
  
end


