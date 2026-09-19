require "test_helper"

class DataTaskControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataTask = dataTasks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataTask" do
    assert_difference("DataTask.count") do
      post dataTasks_url, params: { dataTask: { name:"test string for name", command:"test string for command", retries:100, TaskType:DataTask.TaskTypes[0] } }
    end

    assert_redirected_to dataTasks_url
  end

 
  
  test "should destroy dataTask" do
    assert_difference("DataTask.count", -1) do
      delete dataTask_url(@dataTask)
    end

    assert_redirected_to dataTasks_url
  end
  
end


