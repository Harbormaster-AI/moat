require "test_helper"

class AircraftProgramControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aircraftProgram = aircraftPrograms(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aircraftProgram" do
    assert_difference("AircraftProgram.count") do
      post aircraftPrograms_url, params: { aircraftProgram: { name:"test string for name", programCode:"test string for programCode", entryIntoServiceYear:100, Status:AircraftProgram.Statuss[0] } }
    end

    assert_redirected_to aircraftPrograms_url
  end

 
  
  test "should destroy aircraftProgram" do
    assert_difference("AircraftProgram.count", -1) do
      delete aircraftProgram_url(@aircraftProgram)
    end

    assert_redirected_to aircraftPrograms_url
  end
  
end


