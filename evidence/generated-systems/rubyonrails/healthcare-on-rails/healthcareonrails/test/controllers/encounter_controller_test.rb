require "test_helper"

class EncounterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @encounter = encounters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create encounter" do
    assert_difference("Encounter.count") do
      post encounters_url, params: { encounter: { encounterNumber:"test string for encounterNumber", startDateTime:1.week.ago, endDateTime:1.week.ago, Status:Encounter.Statuss[0], EncounterType:Encounter.EncounterTypes[0] } }
    end

    assert_redirected_to encounters_url
  end

 
  
  test "should destroy encounter" do
    assert_difference("Encounter.count", -1) do
      delete encounter_url(@encounter)
    end

    assert_redirected_to encounters_url
  end
  
end


