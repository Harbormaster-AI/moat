require "test_helper"

class SubrogationRecoveryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @subrogationRecovery = subrogationRecoverys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create subrogationRecovery" do
    assert_difference("SubrogationRecovery.count") do
      post subrogationRecoverys_url, params: { subrogationRecovery: { recoveryReference:"test string for recoveryReference", amount:"test value", recoveryDate:1.week.ago, Status:SubrogationRecovery.Statuss[0] } }
    end

    assert_redirected_to subrogationRecoverys_url
  end

 
  
  test "should destroy subrogationRecovery" do
    assert_difference("SubrogationRecovery.count", -1) do
      delete subrogationRecovery_url(@subrogationRecovery)
    end

    assert_redirected_to subrogationRecoverys_url
  end
  
end


