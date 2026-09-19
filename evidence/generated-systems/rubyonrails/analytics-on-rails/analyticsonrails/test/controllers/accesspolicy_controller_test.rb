require "test_helper"

class AccessPolicyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @accessPolicy = accessPolicys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create accessPolicy" do
    assert_difference("AccessPolicy.count") do
      post accessPolicys_url, params: { accessPolicy: { name:"test string for name", subjectName:"test string for subjectName", AccessLevel:AccessPolicy.AccessLevels[0], SubjectType:AccessPolicy.SubjectTypes[0] } }
    end

    assert_redirected_to accessPolicys_url
  end

 
  
  test "should destroy accessPolicy" do
    assert_difference("AccessPolicy.count", -1) do
      delete accessPolicy_url(@accessPolicy)
    end

    assert_redirected_to accessPolicys_url
  end
  
end


