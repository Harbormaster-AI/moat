require "test_helper"

class AircraftPackageControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftPackage = aircraftPackages(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftPackage" do
    assert_difference("AircraftPackage.count") do
      post aircraftPackages_url, params: { aircraftPackage: { name:"test string for name", PackageType:AircraftPackage.PackageTypes[0] } }
    end

    assert_redirected_to aircraftPackages_url
  end

 
  
  test "should destroy aircraftPackage" do
    assert_difference("AircraftPackage.count", -1) do
      delete aircraftPackage_url(@aircraftPackage)
    end

    assert_redirected_to aircraftPackages_url
  end
  
end


