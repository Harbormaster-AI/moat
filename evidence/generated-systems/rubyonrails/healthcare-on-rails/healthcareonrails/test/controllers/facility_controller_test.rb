require "test_helper"

class FacilityControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @facility = facilitys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create facility" do
    assert_difference("Facility.count") do
      post facilitys_url, params: { facility: { name:"test string for name", facilityCode:"test string for facilityCode", address:"test value", FacilityType:Facility.FacilityTypes[0] } }
    end

    assert_redirected_to facilitys_url
  end

 
  
  test "should destroy facility" do
    assert_difference("Facility.count", -1) do
      delete facility_url(@facility)
    end

    assert_redirected_to facilitys_url
  end
  
end


