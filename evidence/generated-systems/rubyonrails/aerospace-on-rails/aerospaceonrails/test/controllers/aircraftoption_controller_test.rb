require "test_helper"

class AircraftOptionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftOption = aircraftOptions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftOption" do
    assert_difference("AircraftOption.count") do
      post aircraftOptions_url, params: { aircraftOption: { code:"test string for code", name:"test string for name", OptionCategory:AircraftOption.OptionCategorys[0] } }
    end

    assert_redirected_to aircraftOptions_url
  end

 
  
  test "should destroy aircraftOption" do
    assert_difference("AircraftOption.count", -1) do
      delete aircraftOption_url(@aircraftOption)
    end

    assert_redirected_to aircraftOptions_url
  end
  
end


