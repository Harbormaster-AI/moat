require "test_helper"

class WorkCenterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @workCenter = workCenters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create workCenter" do
    assert_difference("WorkCenter.count") do
      post workCenters_url, params: { workCenter: { name:"test string for name", capability:"test string for capability" } }
    end

    assert_redirected_to workCenters_url
  end

 
  
  test "should destroy workCenter" do
    assert_difference("WorkCenter.count", -1) do
      delete workCenter_url(@workCenter)
    end

    assert_redirected_to workCenters_url
  end
  
end


