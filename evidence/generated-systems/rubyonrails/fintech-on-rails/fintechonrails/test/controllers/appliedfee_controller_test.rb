require "test_helper"

class AppliedFeeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @appliedFee = appliedFees(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create appliedFee" do
    assert_difference("AppliedFee.count") do
      post appliedFees_url, params: { appliedFee: { amount:"test value", description:"test string for description", FeeType:AppliedFee.FeeTypes[0] } }
    end

    assert_redirected_to appliedFees_url
  end

 
  
  test "should destroy appliedFee" do
    assert_difference("AppliedFee.count", -1) do
      delete appliedFee_url(@appliedFee)
    end

    assert_redirected_to appliedFees_url
  end
  
end


