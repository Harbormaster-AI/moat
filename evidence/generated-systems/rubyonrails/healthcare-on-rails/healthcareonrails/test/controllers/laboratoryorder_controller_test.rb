require "test_helper"

class LaboratoryOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @laboratoryOrder = laboratoryOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create laboratoryOrder" do
    assert_difference("LaboratoryOrder.count") do
      post laboratoryOrders_url, params: { laboratoryOrder: { testCode:"test string for testCode", fastingRequired:true, SpecimenType:LaboratoryOrder.SpecimenTypes[0] } }
    end

    assert_redirected_to laboratoryOrders_url
  end

 
  
  test "should destroy laboratoryOrder" do
    assert_difference("LaboratoryOrder.count", -1) do
      delete laboratoryOrder_url(@laboratoryOrder)
    end

    assert_redirected_to laboratoryOrders_url
  end
  
end


