require "test_helper"

class AircraftVariantControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftVariant = aircraftVariants(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftVariant" do
    assert_difference("AircraftVariant.count") do
      post aircraftVariants_url, params: { aircraftVariant: { variantCode:"test string for variantCode", rangeNm:100, maxTakeoffWeightKg:"test value" } }
    end

    assert_redirected_to aircraftVariants_url
  end

 
  
  test "should destroy aircraftVariant" do
    assert_difference("AircraftVariant.count", -1) do
      delete aircraftVariant_url(@aircraftVariant)
    end

    assert_redirected_to aircraftVariants_url
  end
  
end


