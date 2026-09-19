require "test_helper"

class MaintenancePlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @maintenancePlan = maintenancePlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create maintenancePlan" do
    assert_difference("MaintenancePlan.count") do
      post maintenancePlans_url, params: { maintenancePlan: { planNumber:"test string for planNumber", interval:1.week.ago, lastServiceDate:1.week.ago, Strategy:MaintenancePlan.Strategys[0] } }
    end

    assert_redirected_to maintenancePlans_url
  end

 
  
  test "should destroy maintenancePlan" do
    assert_difference("MaintenancePlan.count", -1) do
      delete maintenancePlan_url(@maintenancePlan)
    end

    assert_redirected_to maintenancePlans_url
  end
  
end


