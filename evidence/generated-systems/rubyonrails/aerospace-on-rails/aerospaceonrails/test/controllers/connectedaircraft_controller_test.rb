require "test_helper"

class ConnectedAircraftControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @connectedAircraft = connectedAircrafts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create connectedAircraft" do
    assert_difference("ConnectedAircraft.count") do
      post connectedAircrafts_url, params: { connectedAircraft: { communicationsProvider:"test string for communicationsProvider", ConnectivityStatus:ConnectedAircraft.ConnectivityStatuss[0] } }
    end

    assert_redirected_to connectedAircrafts_url
  end

 
  
  test "should destroy connectedAircraft" do
    assert_difference("ConnectedAircraft.count", -1) do
      delete connectedAircraft_url(@connectedAircraft)
    end

    assert_redirected_to connectedAircrafts_url
  end
  
end


