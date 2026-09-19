require "test_helper"

class CabinLayoutControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @cabinLayout = cabinLayouts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create cabinLayout" do
    assert_difference("CabinLayout.count") do
      post cabinLayouts_url, params: { cabinLayout: { layoutCode:"test string for layoutCode", totalSeats:100, classLayout:"test string for classLayout" } }
    end

    assert_redirected_to cabinLayouts_url
  end

 
  
  test "should destroy cabinLayout" do
    assert_difference("CabinLayout.count", -1) do
      delete cabinLayout_url(@cabinLayout)
    end

    assert_redirected_to cabinLayouts_url
  end
  
end


