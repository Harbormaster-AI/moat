require "test_helper"

class MaintenanceOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @maintenanceOrder = maintenanceOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create maintenanceOrder" do
    assert_difference("MaintenanceOrder.count") do
      post maintenanceOrders_url, params: { maintenanceOrder: { orderNumber:"test string for orderNumber", priority:100, requestedDate:1.week.ago, completionDate:1.week.ago, Status:MaintenanceOrder.Statuss[0] } }
    end

    assert_redirected_to maintenanceOrders_url
  end

 
  
  test "should destroy maintenanceOrder" do
    assert_difference("MaintenanceOrder.count", -1) do
      delete maintenanceOrder_url(@maintenanceOrder)
    end

    assert_redirected_to maintenanceOrders_url
  end
  
end


