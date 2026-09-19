require "test_helper"

class SoftwareUpdateControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @softwareUpdate = softwareUpdates(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create softwareUpdate" do
    assert_difference("SoftwareUpdate.count") do
      post softwareUpdates_url, params: { softwareUpdate: { version:"test string for version", appliedDate:1.week.ago, UpdateType:SoftwareUpdate.UpdateTypes[0] } }
    end

    assert_redirected_to softwareUpdates_url
  end

 
  
  test "should destroy softwareUpdate" do
    assert_difference("SoftwareUpdate.count", -1) do
      delete softwareUpdate_url(@softwareUpdate)
    end

    assert_redirected_to softwareUpdates_url
  end
  
end


