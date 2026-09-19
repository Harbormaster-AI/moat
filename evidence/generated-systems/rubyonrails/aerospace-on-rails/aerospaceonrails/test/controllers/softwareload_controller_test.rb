require "test_helper"

class SoftwareLoadControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @softwareLoad = softwareLoads(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create softwareLoad" do
    assert_difference("SoftwareLoad.count") do
      post softwareLoads_url, params: { softwareLoad: { version:"test string for version", LoadType:SoftwareLoad.LoadTypes[0] } }
    end

    assert_redirected_to softwareLoads_url
  end

 
  
  test "should destroy softwareLoad" do
    assert_difference("SoftwareLoad.count", -1) do
      delete softwareLoad_url(@softwareLoad)
    end

    assert_redirected_to softwareLoads_url
  end
  
end


