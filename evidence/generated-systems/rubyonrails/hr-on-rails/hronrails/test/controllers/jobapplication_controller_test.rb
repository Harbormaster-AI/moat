require "test_helper"

class JobApplicationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @jobApplication = jobApplications(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create jobApplication" do
    assert_difference("JobApplication.count") do
      post jobApplications_url, params: { jobApplication: { applicationNumber:"test string for applicationNumber", appliedDate:1.week.ago, resumeUrl:"test string for resumeUrl", Status:JobApplication.Statuss[0] } }
    end

    assert_redirected_to jobApplications_url
  end

 
  
  test "should destroy jobApplication" do
    assert_difference("JobApplication.count", -1) do
      delete jobApplication_url(@jobApplication)
    end

    assert_redirected_to jobApplications_url
  end
  
end


