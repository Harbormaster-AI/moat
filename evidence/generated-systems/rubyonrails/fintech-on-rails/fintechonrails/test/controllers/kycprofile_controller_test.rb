require "test_helper"

class KYCProfileControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @kYCProfile = kYCProfiles(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create kYCProfile" do
    assert_difference("KYCProfile.count") do
      post kYCProfiles_url, params: { kYCProfile: { profileId:"test string for profileId", createdAt:1.week.ago, Status:KYCProfile.Statuss[0], VerificationLevel:KYCProfile.VerificationLevels[0] } }
    end

    assert_redirected_to kYCProfiles_url
  end

 
  
  test "should destroy kYCProfile" do
    assert_difference("KYCProfile.count", -1) do
      delete kYCProfile_url(@kYCProfile)
    end

    assert_redirected_to kYCProfiles_url
  end
  
end


