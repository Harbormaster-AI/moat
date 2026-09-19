require "test_helper"

class TrainingCourseControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @trainingCourse = trainingCourses(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create trainingCourse" do
    assert_difference("TrainingCourse.count") do
      post trainingCourses_url, params: { trainingCourse: { code:"test string for code", title:"test string for title", durationHours:"test value", DeliveryMethod:TrainingCourse.DeliveryMethods[0] } }
    end

    assert_redirected_to trainingCourses_url
  end

 
  
  test "should destroy trainingCourse" do
    assert_difference("TrainingCourse.count", -1) do
      delete trainingCourse_url(@trainingCourse)
    end

    assert_redirected_to trainingCourses_url
  end
  
end


