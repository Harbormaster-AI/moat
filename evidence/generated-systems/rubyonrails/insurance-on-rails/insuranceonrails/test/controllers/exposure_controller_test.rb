require "test_helper"

class ExposureControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @exposure = exposures(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create exposure" do
    assert_difference("Exposure.count") do
      post exposures_url, params: { exposure: { ExposureType:Exposure.ExposureTypes[0], Status:Exposure.Statuss[0] } }
    end

    assert_redirected_to exposures_url
  end

 
  
  test "should destroy exposure" do
    assert_difference("Exposure.count", -1) do
      delete exposure_url(@exposure)
    end

    assert_redirected_to exposures_url
  end
  
end


