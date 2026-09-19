require "test_helper"

class LandingGearControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @landingGear = landingGears(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create landingGear" do
    assert_difference("LandingGear.count") do
      post landingGears_url, params: { landingGear: { supplierPartNumber:"test string for supplierPartNumber", GearType:LandingGear.GearTypes[0] } }
    end

    assert_redirected_to landingGears_url
  end

 
  
  test "should destroy landingGear" do
    assert_difference("LandingGear.count", -1) do
      delete landingGear_url(@landingGear)
    end

    assert_redirected_to landingGears_url
  end
  
end


