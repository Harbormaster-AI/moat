require "test_helper"

class ClinicalOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @clinicalOrder = clinicalOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create clinicalOrder" do
    assert_difference("ClinicalOrder.count") do
      post clinicalOrders_url, params: { clinicalOrder: { orderNumber:"test string for orderNumber", Status:ClinicalOrder.Statuss[0], OrderType:ClinicalOrder.OrderTypes[0], Priority:ClinicalOrder.Prioritys[0] } }
    end

    assert_redirected_to clinicalOrders_url
  end

 
  
  test "should destroy clinicalOrder" do
    assert_difference("ClinicalOrder.count", -1) do
      delete clinicalOrder_url(@clinicalOrder)
    end

    assert_redirected_to clinicalOrders_url
  end
  
end


