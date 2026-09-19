require "test_helper"

class ExpirationPolicyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @expirationPolicy = expirationPolicys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create expirationPolicy" do
    assert_difference("ExpirationPolicy.count") do
      post expirationPolicys_url, params: { expirationPolicy: { rejectIfDaysToExpireLessThan:100, autoQuarantineDaysToExpire:100, RotationMethod:ExpirationPolicy.RotationMethods[0] } }
    end

    assert_redirected_to expirationPolicys_url
  end

 
  
  test "should destroy expirationPolicy" do
    assert_difference("ExpirationPolicy.count", -1) do
      delete expirationPolicy_url(@expirationPolicy)
    end

    assert_redirected_to expirationPolicys_url
  end
  
end


