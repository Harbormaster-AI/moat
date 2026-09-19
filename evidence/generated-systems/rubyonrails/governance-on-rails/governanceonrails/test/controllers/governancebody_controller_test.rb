require "test_helper"

class GovernanceBodyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @governanceBody = governanceBodys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create governanceBody" do
    assert_difference("GovernanceBody.count") do
      post governanceBodys_url, params: { governanceBody: { name:"test string for name", charterUrl:"test value", chair:"test string for chair", BodyType:GovernanceBody.BodyTypes[0] } }
    end

    assert_redirected_to governanceBodys_url
  end

 
  
  test "should destroy governanceBody" do
    assert_difference("GovernanceBody.count", -1) do
      delete governanceBody_url(@governanceBody)
    end

    assert_redirected_to governanceBodys_url
  end
  
end


