require "test_helper"

class InventoryItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inventoryItem = inventoryItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inventoryItem" do
    assert_difference("InventoryItem.count") do
      post inventoryItems_url, params: { inventoryItem: { quantityOnHand:"test value", quantityReserved:"test value", lotNumber:"test value", serialNumber:"test value" } }
    end

    assert_redirected_to inventoryItems_url
  end

 
  
  test "should destroy inventoryItem" do
    assert_difference("InventoryItem.count", -1) do
      delete inventoryItem_url(@inventoryItem)
    end

    assert_redirected_to inventoryItems_url
  end
  
end


