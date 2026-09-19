require "test_helper"

class AircraftFamilyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftFamily = aircraftFamilys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftFamily" do
    assert_difference("AircraftFamily.count") do
      post aircraftFamilys_url, params: { aircraftFamily: { name:"test string for name", familyCode:"test string for familyCode" } }
    end

    assert_redirected_to aircraftFamilys_url
  end

 
  
  test "should destroy aircraftFamily" do
    assert_difference("AircraftFamily.count", -1) do
      delete aircraftFamily_url(@aircraftFamily)
    end

    assert_redirected_to aircraftFamilys_url
  end
  
end


