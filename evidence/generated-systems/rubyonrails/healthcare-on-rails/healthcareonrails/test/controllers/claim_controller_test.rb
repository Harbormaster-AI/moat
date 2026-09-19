require "test_helper"

class ClaimControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @claim = claims(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create claim" do
    assert_difference("Claim.count") do
      post claims_url, params: { claim: { claimNumber:"test string for claimNumber", totalAmount:"test value", Status:Claim.Statuss[0] } }
    end

    assert_redirected_to claims_url
  end

 
  
  test "should destroy claim" do
    assert_difference("Claim.count", -1) do
      delete claim_url(@claim)
    end

    assert_redirected_to claims_url
  end
  
end


