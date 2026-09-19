require "test_helper"

class DataBreachControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataBreach = dataBreachs(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataBreach" do
    assert_difference("DataBreach.count") do
      post dataBreachs_url, params: { dataBreach: { incidentDate:1.week.ago, description:"test string for description", recordsAffected:100, notificationRequired:true, Severity:DataBreach.Severitys[0], Status:DataBreach.Statuss[0] } }
    end

    assert_redirected_to dataBreachs_url
  end

 
  
  test "should destroy dataBreach" do
    assert_difference("DataBreach.count", -1) do
      delete dataBreach_url(@dataBreach)
    end

    assert_redirected_to dataBreachs_url
  end
  
end


