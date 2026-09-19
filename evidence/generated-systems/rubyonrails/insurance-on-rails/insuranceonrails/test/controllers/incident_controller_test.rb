require "test_helper"

class IncidentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @incident = incidents(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create incident" do
    assert_difference("Incident.count") do
      post incidents_url, params: { incident: { location:"test value", description:"test string for description", IncidentType:Incident.IncidentTypes[0] } }
    end

    assert_redirected_to incidents_url
  end

 
  
  test "should destroy incident" do
    assert_difference("Incident.count", -1) do
      delete incident_url(@incident)
    end

    assert_redirected_to incidents_url
  end
  
end


