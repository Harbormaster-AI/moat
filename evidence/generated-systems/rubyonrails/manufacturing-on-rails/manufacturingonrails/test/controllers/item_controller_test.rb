require "test_helper"

class ItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @item = items(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create item" do
    assert_difference("Item.count") do
      post items_url, params: { item: { itemNumber:"test string for itemNumber", name:"test string for name", standardCost:"test value", weight:"test value", asSerialControlled:true, ItemType:Item.ItemTypes[0], ProcurementType:Item.ProcurementTypes[0], UnitOfMeasure:Item.UnitOfMeasures[0], LifecycleStatus:Item.LifecycleStatuss[0] } }
    end

    assert_redirected_to items_url
  end

 
  
  test "should destroy item" do
    assert_difference("Item.count", -1) do
      delete item_url(@item)
    end

    assert_redirected_to items_url
  end
  
end


