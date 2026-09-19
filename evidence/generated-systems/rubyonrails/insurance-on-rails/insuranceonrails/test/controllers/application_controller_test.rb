require "test_helper"

class ApplicationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @application = applications(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create application" do
    assert_difference("Application.count") do
      post applications_url, params: { application: { applicationNumber:"test string for applicationNumber", submissionDate:1.week.ago, Status:Application.Statuss[0] } }
    end

    assert_redirected_to applications_url
  end

 
  
  test "should destroy application" do
    assert_difference("Application.count", -1) do
      delete application_url(@application)
    end

    assert_redirected_to applications_url
  end
  
end


