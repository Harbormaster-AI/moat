require "test_helper"

class AircraftModelControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftModel = aircraftModels(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftModel" do
    assert_difference("AircraftModel.count") do
      post aircraftModels_url, params: { aircraftModel: { name:"test string for name", modelDesignation:"test string for modelDesignation", AircraftType:AircraftModel.AircraftTypes[0] } }
    end

    assert_redirected_to aircraftModels_url
  end

 
  
  test "should destroy aircraftModel" do
    assert_difference("AircraftModel.count", -1) do
      delete aircraftModel_url(@aircraftModel)
    end

    assert_redirected_to aircraftModels_url
  end
  
end


