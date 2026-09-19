require "test_helper"

class CycleCountEntryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @cycleCountEntry = cycleCountEntrys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create cycleCountEntry" do
    assert_difference("CycleCountEntry.count") do
      post cycleCountEntrys_url, params: { cycleCountEntry: { lineNumber:100, systemQuantity:"test value", countedQuantity:"test value", varianceQuantity:"test value", recountRequired:true, StockStatus:CycleCountEntry.StockStatuss[0] } }
    end

    assert_redirected_to cycleCountEntrys_url
  end

 
  
  test "should destroy cycleCountEntry" do
    assert_difference("CycleCountEntry.count", -1) do
      delete cycleCountEntry_url(@cycleCountEntry)
    end

    assert_redirected_to cycleCountEntrys_url
  end
  
end


