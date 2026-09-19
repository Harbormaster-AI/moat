require "test_helper"

class LocationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @location = locations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create location" do
    assert_difference("Location.count") do
      post locations_url, params: { location: { name:"test string for name", address:"test value", timezone:"test string for timezone" } }
    end

    assert_redirected_to locations_url
  end

 
  
  test "should destroy location" do
    assert_difference("Location.count", -1) do
      delete location_url(@location)
    end

    assert_redirected_to locations_url
  end
  
end


