require "test_helper"

class AuditEngagementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @auditEngagement = auditEngagements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create auditEngagement" do
    assert_difference("AuditEngagement.count") do
      post auditEngagements_url, params: { auditEngagement: { title:"test string for title", startDate:1.week.ago, endDate:1.week.ago, Status:AuditEngagement.Statuss[0] } }
    end

    assert_redirected_to auditEngagements_url
  end

 
  
  test "should destroy auditEngagement" do
    assert_difference("AuditEngagement.count", -1) do
      delete auditEngagement_url(@auditEngagement)
    end

    assert_redirected_to auditEngagements_url
  end
  
end


