require "test_helper"

class AuditFindingControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @auditFinding = auditFindings(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create auditFinding" do
    assert_difference("AuditFinding.count") do
      post auditFindings_url, params: { auditFinding: { title:"test string for title", description:"test string for description", dueDate:1.week.ago, Severity:AuditFinding.Severitys[0], Status:AuditFinding.Statuss[0] } }
    end

    assert_redirected_to auditFindings_url
  end

 
  
  test "should destroy auditFinding" do
    assert_difference("AuditFinding.count", -1) do
      delete auditFinding_url(@auditFinding)
    end

    assert_redirected_to auditFindings_url
  end
  
end


