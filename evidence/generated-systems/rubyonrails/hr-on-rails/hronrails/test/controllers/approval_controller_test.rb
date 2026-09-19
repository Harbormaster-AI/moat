require "test_helper"

class ApprovalControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @approval = approvals(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create approval" do
    assert_difference("Approval.count") do
      post approvals_url, params: { approval: { approverComment:"test string for approverComment", actionDate:1.week.ago, Status:Approval.Statuss[0] } }
    end

    assert_redirected_to approvals_url
  end

 
  
  test "should destroy approval" do
    assert_difference("Approval.count", -1) do
      delete approval_url(@approval)
    end

    assert_redirected_to approvals_url
  end
  
end


