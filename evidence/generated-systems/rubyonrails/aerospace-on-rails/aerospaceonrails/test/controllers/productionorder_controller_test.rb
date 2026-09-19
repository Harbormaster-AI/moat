require "test_helper"

class ProductionOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productionOrder = productionOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productionOrder" do
    assert_difference("ProductionOrder.count") do
      post productionOrders_url, params: { productionOrder: { orderNumber:"test string for orderNumber", Status:ProductionOrder.Statuss[0] } }
    end

    assert_redirected_to productionOrders_url
  end

 
  
  test "should destroy productionOrder" do
    assert_difference("ProductionOrder.count", -1) do
      delete productionOrder_url(@productionOrder)
    end

    assert_redirected_to productionOrders_url
  end
  
end


