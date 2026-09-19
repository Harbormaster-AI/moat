require "test_helper"

class Exception_ControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @exception_ = exception_s(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create exception_" do
    assert_difference("Exception_.count") do
      post exception_s_url, params: { exception_: { title:"test string for title", justification:"test string for justification", startDate:1.week.ago, endDate:1.week.ago, ExceptionType:Exception_.ExceptionTypes[0], Status:Exception_.Statuss[0] } }
    end

    assert_redirected_to exception_s_url
  end

 
  
  test "should destroy exception_" do
    assert_difference("Exception_.count", -1) do
      delete exception__url(@exception_)
    end

    assert_redirected_to exception_s_url
  end
  
end


