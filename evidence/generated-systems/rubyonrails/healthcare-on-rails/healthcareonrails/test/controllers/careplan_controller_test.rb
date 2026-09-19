require "test_helper"

class CarePlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @carePlan = carePlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create carePlan" do
    assert_difference("CarePlan.count") do
      post carePlans_url, params: { carePlan: { planNumber:"test string for planNumber", goalSummary:"test string for goalSummary", Status:CarePlan.Statuss[0] } }
    end

    assert_redirected_to carePlans_url
  end

 
  
  test "should destroy carePlan" do
    assert_difference("CarePlan.count", -1) do
      delete carePlan_url(@carePlan)
    end

    assert_redirected_to carePlans_url
  end
  
end


