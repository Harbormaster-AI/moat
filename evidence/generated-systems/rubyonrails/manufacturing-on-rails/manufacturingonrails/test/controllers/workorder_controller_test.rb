require "test_helper"

class WorkOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @workOrder = workOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create workOrder" do
    assert_difference("WorkOrder.count") do
      post workOrders_url, params: { workOrder: { workOrderNumber:"test string for workOrderNumber", plannedStart:1.week.ago, plannedEnd:1.week.ago, quantity:"test value", priority:100, Status:WorkOrder.Statuss[0] } }
    end

    assert_redirected_to workOrders_url
  end

 
  
  test "should destroy workOrder" do
    assert_difference("WorkOrder.count", -1) do
      delete workOrder_url(@workOrder)
    end

    assert_redirected_to workOrders_url
  end
  
end


