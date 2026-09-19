require "test_helper"

class ActivityControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @activity = activitys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create activity" do
    assert_difference("Activity.count") do
      post activitys_url, params: { activity: { subject:"test string for subject", dueDate:1.week.ago, startAt:1.week.ago, endAt:1.week.ago, location:"test string for location", ActivityType:Activity.ActivityTypes[0], Status:Activity.Statuss[0], Priority:Activity.Prioritys[0] } }
    end

    assert_redirected_to activitys_url
  end

 
  
  test "should destroy activity" do
    assert_difference("Activity.count", -1) do
      delete activity_url(@activity)
    end

    assert_redirected_to activitys_url
  end
  
end


