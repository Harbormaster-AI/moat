require "test_helper"

class InspectionPlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inspectionPlan = inspectionPlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inspectionPlan" do
    assert_difference("InspectionPlan.count") do
      post inspectionPlans_url, params: { inspectionPlan: { planNumber:"test string for planNumber", revision:"test string for revision", SamplingPlan:InspectionPlan.SamplingPlans[0], Status:InspectionPlan.Statuss[0] } }
    end

    assert_redirected_to inspectionPlans_url
  end

 
  
  test "should destroy inspectionPlan" do
    assert_difference("InspectionPlan.count", -1) do
      delete inspectionPlan_url(@inspectionPlan)
    end

    assert_redirected_to inspectionPlans_url
  end
  
end


