require "test_helper"

class StockKeepingUnitControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @stockKeepingUnit = stockKeepingUnits(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create stockKeepingUnit" do
    assert_difference("StockKeepingUnit.count") do
      post stockKeepingUnits_url, params: { stockKeepingUnit: { skuCode:"test value", name:"test string for name", weight:"test value", weightUnit:"test string for weightUnit", volume:"test value", volumeUnit:"test string for volumeUnit", shelfLifeDays:100, hazardousMaterial:true, ItemType:StockKeepingUnit.ItemTypes[0], UnitOfMeasure:StockKeepingUnit.UnitOfMeasures[0] } }
    end

    assert_redirected_to stockKeepingUnits_url
  end

 
  
  test "should destroy stockKeepingUnit" do
    assert_difference("StockKeepingUnit.count", -1) do
      delete stockKeepingUnit_url(@stockKeepingUnit)
    end

    assert_redirected_to stockKeepingUnits_url
  end
  
end


