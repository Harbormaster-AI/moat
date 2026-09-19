require "test_helper"

class AnalyticsWorkspaceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @analyticsWorkspace = analyticsWorkspaces(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create analyticsWorkspace" do
    assert_difference("AnalyticsWorkspace.count") do
      post analyticsWorkspaces_url, params: { analyticsWorkspace: { name:"test string for name", businessDomain:"test string for businessDomain", ownerTeam:"test string for ownerTeam", GovernanceTier:AnalyticsWorkspace.GovernanceTiers[0] } }
    end

    assert_redirected_to analyticsWorkspaces_url
  end

 
  
  test "should destroy analyticsWorkspace" do
    assert_difference("AnalyticsWorkspace.count", -1) do
      delete analyticsWorkspace_url(@analyticsWorkspace)
    end

    assert_redirected_to analyticsWorkspaces_url
  end
  
end


