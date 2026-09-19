require "test_helper"

class VerifiedAddressControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @verifiedAddress = verifiedAddresss(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create verifiedAddress" do
    assert_difference("VerifiedAddress.count") do
      post verifiedAddresss_url, params: { verifiedAddress: { address:"test value", verifiedAt:1.week.ago, VerificationStatus:VerifiedAddress.VerificationStatuss[0] } }
    end

    assert_redirected_to verifiedAddresss_url
  end

 
  
  test "should destroy verifiedAddress" do
    assert_difference("VerifiedAddress.count", -1) do
      delete verifiedAddress_url(@verifiedAddress)
    end

    assert_redirected_to verifiedAddresss_url
  end
  
end


