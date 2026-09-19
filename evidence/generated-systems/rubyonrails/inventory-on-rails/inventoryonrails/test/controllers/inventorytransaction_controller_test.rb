require "test_helper"

class InventoryTransactionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inventoryTransaction = inventoryTransactions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inventoryTransaction" do
    assert_difference("InventoryTransaction.count") do
      post inventoryTransactions_url, params: { inventoryTransaction: { transactionNumber:"test string for transactionNumber", quantity:"test value", unitCost:"test value", transactionDate:1.week.ago, reasonCode:"test string for reasonCode", TransactionType:InventoryTransaction.TransactionTypes[0], UnitOfMeasure:InventoryTransaction.UnitOfMeasures[0], Status:InventoryTransaction.Statuss[0] } }
    end

    assert_redirected_to inventoryTransactions_url
  end

 
  
  test "should destroy inventoryTransaction" do
    assert_difference("InventoryTransaction.count", -1) do
      delete inventoryTransaction_url(@inventoryTransaction)
    end

    assert_redirected_to inventoryTransactions_url
  end
  
end


