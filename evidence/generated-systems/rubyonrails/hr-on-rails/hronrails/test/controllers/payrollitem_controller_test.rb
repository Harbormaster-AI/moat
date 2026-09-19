require "test_helper"

class PayrollItemControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @payrollItem = payrollItems(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create payrollItem" do
    assert_difference("PayrollItem.count") do
      post payrollItems_url, params: { payrollItem: { amount:"test value", taxable:true, ItemType:PayrollItem.ItemTypes[0] } }
    end

    assert_redirected_to payrollItems_url
  end

 
  
  test "should destroy payrollItem" do
    assert_difference("PayrollItem.count", -1) do
      delete payrollItem_url(@payrollItem)
    end

    assert_redirected_to payrollItems_url
  end
  
end


