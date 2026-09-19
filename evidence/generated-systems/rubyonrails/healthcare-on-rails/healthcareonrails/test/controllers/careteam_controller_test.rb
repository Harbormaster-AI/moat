require "test_helper"

class CareTeamControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @careTeam = careTeams(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create careTeam" do
    assert_difference("CareTeam.count") do
      post careTeams_url, params: { careTeam: { name:"test string for name", CareSetting:CareTeam.CareSettings[0] } }
    end

    assert_redirected_to careTeams_url
  end

 
  
  test "should destroy careTeam" do
    assert_difference("CareTeam.count", -1) do
      delete careTeam_url(@careTeam)
    end

    assert_redirected_to careTeams_url
  end
  
end


