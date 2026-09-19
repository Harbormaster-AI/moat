require "test_helper"

class Case_ControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @case_ = case_s(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create case_" do
    assert_difference("Case_.count") do
      post case_s_url, params: { case_: { caseNumber:"test string for caseNumber", subject:"test string for subject", description:"test string for description", slaDue:1.week.ago, Status:Case_.Statuss[0], Priority:Case_.Prioritys[0], Origin:Case_.Origins[0], Severity:Case_.Severitys[0] } }
    end

    assert_redirected_to case_s_url
  end

 
  
  test "should destroy case_" do
    assert_difference("Case_.count", -1) do
      delete case__url(@case_)
    end

    assert_redirected_to case_s_url
  end
  
end


