require "test_helper"

class TrainingEnrollmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @trainingEnrollment = trainingEnrollments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create trainingEnrollment" do
    assert_difference("TrainingEnrollment.count") do
      post trainingEnrollments_url, params: { trainingEnrollment: { enrollmentNumber:"test string for enrollmentNumber", completionDate:1.week.ago, score:"test value", Status:TrainingEnrollment.Statuss[0] } }
    end

    assert_redirected_to trainingEnrollments_url
  end

 
  
  test "should destroy trainingEnrollment" do
    assert_difference("TrainingEnrollment.count", -1) do
      delete trainingEnrollment_url(@trainingEnrollment)
    end

    assert_redirected_to trainingEnrollments_url
  end
  
end


