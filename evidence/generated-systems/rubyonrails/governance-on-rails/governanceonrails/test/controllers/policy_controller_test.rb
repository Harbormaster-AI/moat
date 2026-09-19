require "test_helper"

class PolicyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @policy = policys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create policy" do
    assert_difference("Policy.count") do
      post policys_url, params: { policy: { title:"test string for title", versionLabel:"test string for versionLabel", approvalDate:1.week.ago, nextReviewDate:1.week.ago, documentUrl:"test value", PolicyType:Policy.PolicyTypes[0], Status:Policy.Statuss[0] } }
    end

    assert_redirected_to policys_url
  end

 
  
  test "should destroy policy" do
    assert_difference("Policy.count", -1) do
      delete policy_url(@policy)
    end

    assert_redirected_to policys_url
  end
  
end


