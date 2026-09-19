require "test_helper"

class ImagingCenterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @imagingCenter = imagingCenters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create imagingCenter" do
    assert_difference("ImagingCenter.count") do
      post imagingCenters_url, params: { imagingCenter: { name:"test string for name" } }
    end

    assert_redirected_to imagingCenters_url
  end

 
  
  test "should destroy imagingCenter" do
    assert_difference("ImagingCenter.count", -1) do
      delete imagingCenter_url(@imagingCenter)
    end

    assert_redirected_to imagingCenters_url
  end
  
end


