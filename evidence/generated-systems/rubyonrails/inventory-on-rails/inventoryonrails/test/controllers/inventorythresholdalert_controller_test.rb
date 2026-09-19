require "test_helper"

class InventoryThresholdAlertControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inventoryThresholdAlert = inventoryThresholdAlerts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inventoryThresholdAlert" do
    assert_difference("InventoryThresholdAlert.count") do
      post inventoryThresholdAlerts_url, params: { inventoryThresholdAlert: { alertNumber:"test string for alertNumber", detectedAt:1.week.ago, message:"test string for message", AlertType:InventoryThresholdAlert.AlertTypes[0], Status:InventoryThresholdAlert.Statuss[0] } }
    end

    assert_redirected_to inventoryThresholdAlerts_url
  end

 
  
  test "should destroy inventoryThresholdAlert" do
    assert_difference("InventoryThresholdAlert.count", -1) do
      delete inventoryThresholdAlert_url(@inventoryThresholdAlert)
    end

    assert_redirected_to inventoryThresholdAlerts_url
  end
  
end


