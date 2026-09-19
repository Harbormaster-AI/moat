require "test_helper"

class MaintenanceWorkOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @maintenanceWorkOrder = maintenanceWorkOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create maintenanceWorkOrder" do
    assert_difference("MaintenanceWorkOrder.count") do
      post maintenanceWorkOrders_url, params: { maintenanceWorkOrder: { workOrderNumber:"test string for workOrderNumber", Status:MaintenanceWorkOrder.Statuss[0] } }
    end

    assert_redirected_to maintenanceWorkOrders_url
  end

 
  
  test "should destroy maintenanceWorkOrder" do
    assert_difference("MaintenanceWorkOrder.count", -1) do
      delete maintenanceWorkOrder_url(@maintenanceWorkOrder)
    end

    assert_redirected_to maintenanceWorkOrders_url
  end
  
end


