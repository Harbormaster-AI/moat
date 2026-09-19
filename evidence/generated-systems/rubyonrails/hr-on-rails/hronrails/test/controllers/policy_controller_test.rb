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
      post policys_url, params: { policy: { policyNumber:"test string for policyNumber", name:"test string for name", effectiveDate:1.week.ago, description:"test string for description" } }
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


