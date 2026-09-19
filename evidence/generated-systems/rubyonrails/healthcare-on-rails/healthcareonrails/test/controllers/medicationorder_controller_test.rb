require "test_helper"

class MedicationOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @medicationOrder = medicationOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create medicationOrder" do
    assert_difference("MedicationOrder.count") do
      post medicationOrders_url, params: { medicationOrder: { medicationCode:"test string for medicationCode", dose:"test value", frequency:"test string for frequency", duration:"test string for duration", Route:MedicationOrder.Routes[0] } }
    end

    assert_redirected_to medicationOrders_url
  end

 
  
  test "should destroy medicationOrder" do
    assert_difference("MedicationOrder.count", -1) do
      delete medicationOrder_url(@medicationOrder)
    end

    assert_redirected_to medicationOrders_url
  end
  
end


