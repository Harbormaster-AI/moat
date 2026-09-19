require "test_helper"

class JobProfileControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @jobProfile = jobProfiles(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create jobProfile" do
    assert_difference("JobProfile.count") do
      post jobProfiles_url, params: { jobProfile: { title:"test string for title", jobCode:"test string for jobCode", JobLevel:JobProfile.JobLevels[0], ExemptStatus:JobProfile.ExemptStatuss[0] } }
    end

    assert_redirected_to jobProfiles_url
  end

 
  
  test "should destroy jobProfile" do
    assert_difference("JobProfile.count", -1) do
      delete jobProfile_url(@jobProfile)
    end

    assert_redirected_to jobProfiles_url
  end
  
end


