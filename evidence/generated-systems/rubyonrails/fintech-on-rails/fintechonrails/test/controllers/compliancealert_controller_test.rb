require "test_helper"

class ComplianceAlertControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @complianceAlert = complianceAlerts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create complianceAlert" do
    assert_difference("ComplianceAlert.count") do
      post complianceAlerts_url, params: { complianceAlert: { alertCode:"test string for alertCode", raisedAt:1.week.ago, notes:"test string for notes", Severity:ComplianceAlert.Severitys[0], Status:ComplianceAlert.Statuss[0] } }
    end

    assert_redirected_to complianceAlerts_url
  end

 
  
  test "should destroy complianceAlert" do
    assert_difference("ComplianceAlert.count", -1) do
      delete complianceAlert_url(@complianceAlert)
    end

    assert_redirected_to complianceAlerts_url
  end
  
end


