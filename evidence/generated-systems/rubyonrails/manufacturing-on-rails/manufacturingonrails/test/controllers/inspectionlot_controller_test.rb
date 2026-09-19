require "test_helper"

class InspectionLotControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inspectionLot = inspectionLots(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inspectionLot" do
    assert_difference("InspectionLot.count") do
      post inspectionLots_url, params: { inspectionLot: { lotNumber:"test string for lotNumber", quantity:"test value", sampleSize:100, createdOn:1.week.ago, InspectionType:InspectionLot.InspectionTypes[0], Status:InspectionLot.Statuss[0] } }
    end

    assert_redirected_to inspectionLots_url
  end

 
  
  test "should destroy inspectionLot" do
    assert_difference("InspectionLot.count", -1) do
      delete inspectionLot_url(@inspectionLot)
    end

    assert_redirected_to inspectionLots_url
  end
  
end


