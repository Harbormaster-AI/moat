require "test_helper"

class GoalControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @goal = goals(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create goal" do
    assert_difference("Goal.count") do
      post goals_url, params: { goal: { title:"test string for title", description:"test string for description", targetDate:1.week.ago, weight:"test value", Status:Goal.Statuss[0] } }
    end

    assert_redirected_to goals_url
  end

 
  
  test "should destroy goal" do
    assert_difference("Goal.count", -1) do
      delete goal_url(@goal)
    end

    assert_redirected_to goals_url
  end
  
end


