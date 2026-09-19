require "test_helper"

class CorrectiveActionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @correctiveAction = correctiveActions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create correctiveAction" do
    assert_difference("CorrectiveAction.count") do
      post correctiveActions_url, params: { correctiveAction: { actionTitle:"test string for actionTitle", owner:"test string for owner", targetDate:1.week.ago, Status:CorrectiveAction.Statuss[0] } }
    end

    assert_redirected_to correctiveActions_url
  end

 
  
  test "should destroy correctiveAction" do
    assert_difference("CorrectiveAction.count", -1) do
      delete correctiveAction_url(@correctiveAction)
    end

    assert_redirected_to correctiveActions_url
  end
  
end


