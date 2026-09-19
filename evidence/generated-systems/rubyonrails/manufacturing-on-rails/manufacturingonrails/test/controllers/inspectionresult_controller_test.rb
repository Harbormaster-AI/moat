require "test_helper"

class InspectionResultControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inspectionResult = inspectionResults(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inspectionResult" do
    assert_difference("InspectionResult.count") do
      post inspectionResults_url, params: { inspectionResult: { resultValue:"test value", recordedOn:1.week.ago, notes:"test string for notes", ResultStatus:InspectionResult.ResultStatuss[0] } }
    end

    assert_redirected_to inspectionResults_url
  end

 
  
  test "should destroy inspectionResult" do
    assert_difference("InspectionResult.count", -1) do
      delete inspectionResult_url(@inspectionResult)
    end

    assert_redirected_to inspectionResults_url
  end
  
end


