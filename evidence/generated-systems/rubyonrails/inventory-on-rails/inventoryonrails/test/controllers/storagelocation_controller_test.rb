require "test_helper"

class StorageLocationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @storageLocation = storageLocations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create storageLocation" do
    assert_difference("StorageLocation.count") do
      post storageLocations_url, params: { storageLocation: { code:"test string for code", temperatureControlled:true, capacity:"test value", capacityUnit:"test string for capacityUnit", LocationType:StorageLocation.LocationTypes[0] } }
    end

    assert_redirected_to storageLocations_url
  end

 
  
  test "should destroy storageLocation" do
    assert_difference("StorageLocation.count", -1) do
      delete storageLocation_url(@storageLocation)
    end

    assert_redirected_to storageLocations_url
  end
  
end


