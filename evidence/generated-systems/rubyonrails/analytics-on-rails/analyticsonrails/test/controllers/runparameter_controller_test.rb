require "test_helper"

class RunParameterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @runParameter = runParameters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create runParameter" do
    assert_difference("RunParameter.count") do
      post runParameters_url, params: { runParameter: { name:"test string for name", value:"test string for value" } }
    end

    assert_redirected_to runParameters_url
  end

 
  
  test "should destroy runParameter" do
    assert_difference("RunParameter.count", -1) do
      delete runParameter_url(@runParameter)
    end

    assert_redirected_to runParameters_url
  end
  
end


