require "test_helper"

class InterviewControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @interview = interviews(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create interview" do
    assert_difference("Interview.count") do
      post interviews_url, params: { interview: { interviewDate:1.week.ago, feedback:"test string for feedback", Stage:Interview.Stages[0], Result:Interview.Results[0] } }
    end

    assert_redirected_to interviews_url
  end

 
  
  test "should destroy interview" do
    assert_difference("Interview.count", -1) do
      delete interview_url(@interview)
    end

    assert_redirected_to interviews_url
  end
  
end


