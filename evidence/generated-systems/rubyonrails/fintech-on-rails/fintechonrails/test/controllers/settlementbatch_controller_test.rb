require "test_helper"

class SettlementBatchControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @settlementBatch = settlementBatchs(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create settlementBatch" do
    assert_difference("SettlementBatch.count") do
      post settlementBatchs_url, params: { settlementBatch: { batchId:"test string for batchId", periodStart:1.week.ago, periodEnd:1.week.ago, totalVolume:"test value", totalCount:100, Status:SettlementBatch.Statuss[0] } }
    end

    assert_redirected_to settlementBatchs_url
  end

 
  
  test "should destroy settlementBatch" do
    assert_difference("SettlementBatch.count", -1) do
      delete settlementBatch_url(@settlementBatch)
    end

    assert_redirected_to settlementBatchs_url
  end
  
end


