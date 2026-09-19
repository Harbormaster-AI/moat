require "test_helper"

class MedicalDeviceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @medicalDevice = medicalDevices(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create medicalDevice" do
    assert_difference("MedicalDevice.count") do
      post medicalDevices_url, params: { medicalDevice: { udi:"test string for udi", manufacturer:"test string for manufacturer", DeviceType:MedicalDevice.DeviceTypes[0], ConnectivityStatus:MedicalDevice.ConnectivityStatuss[0] } }
    end

    assert_redirected_to medicalDevices_url
  end

 
  
  test "should destroy medicalDevice" do
    assert_difference("MedicalDevice.count", -1) do
      delete medicalDevice_url(@medicalDevice)
    end

    assert_redirected_to medicalDevices_url
  end
  
end


