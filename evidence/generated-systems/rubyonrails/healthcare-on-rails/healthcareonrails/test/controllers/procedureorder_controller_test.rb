require "test_helper"

class ProcedureOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @procedureOrder = procedureOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create procedureOrder" do
    assert_difference("ProcedureOrder.count") do
      post procedureOrders_url, params: { procedureOrder: { procedureCode:"test string for procedureCode", consentObtained:true, AnesthesiaType:ProcedureOrder.AnesthesiaTypes[0] } }
    end

    assert_redirected_to procedureOrders_url
  end

 
  
  test "should destroy procedureOrder" do
    assert_difference("ProcedureOrder.count", -1) do
      delete procedureOrder_url(@procedureOrder)
    end

    assert_redirected_to procedureOrders_url
  end
  
end


