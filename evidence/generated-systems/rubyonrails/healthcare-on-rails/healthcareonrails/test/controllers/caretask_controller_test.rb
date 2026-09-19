require "test_helper"

class CareTaskControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @careTask = careTasks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create careTask" do
    assert_difference("CareTask.count") do
      post careTasks_url, params: { careTask: { description:"test string for description", dueDate:1.week.ago, Status:CareTask.Statuss[0], Priority:CareTask.Prioritys[0] } }
    end

    assert_redirected_to careTasks_url
  end

 
  
  test "should destroy careTask" do
    assert_difference("CareTask.count", -1) do
      delete careTask_url(@careTask)
    end

    assert_redirected_to careTasks_url
  end
  
end


