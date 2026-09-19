require "test_helper"

class ImagingOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @imagingOrder = imagingOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create imagingOrder" do
    assert_difference("ImagingOrder.count") do
      post imagingOrders_url, params: { imagingOrder: { bodySite:"test string for bodySite", contrast:true, Modality:ImagingOrder.Modalitys[0] } }
    end

    assert_redirected_to imagingOrders_url
  end

 
  
  test "should destroy imagingOrder" do
    assert_difference("ImagingOrder.count", -1) do
      delete imagingOrder_url(@imagingOrder)
    end

    assert_redirected_to imagingOrders_url
  end
  
end


