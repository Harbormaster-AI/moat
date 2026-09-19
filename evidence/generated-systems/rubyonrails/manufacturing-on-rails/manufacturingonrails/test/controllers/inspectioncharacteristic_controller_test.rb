require "test_helper"

class InspectionCharacteristicControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inspectionCharacteristic = inspectionCharacteristics(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inspectionCharacteristic" do
    assert_difference("InspectionCharacteristic.count") do
      post inspectionCharacteristics_url, params: { inspectionCharacteristic: { characteristicCode:"test string for characteristicCode", name:"test string for name", lowerSpecLimit:"test value", upperSpecLimit:"test value", target:"test value", MeasurementType:InspectionCharacteristic.MeasurementTypes[0] } }
    end

    assert_redirected_to inspectionCharacteristics_url
  end

 
  
  test "should destroy inspectionCharacteristic" do
    assert_difference("InspectionCharacteristic.count", -1) do
      delete inspectionCharacteristic_url(@inspectionCharacteristic)
    end

    assert_redirected_to inspectionCharacteristics_url
  end
  
end


