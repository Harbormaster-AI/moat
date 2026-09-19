require "test_helper"

class JobFamilyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @jobFamily = jobFamilys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create jobFamily" do
    assert_difference("JobFamily.count") do
      post jobFamilys_url, params: { jobFamily: { name:"test string for name", description:"test string for description" } }
    end

    assert_redirected_to jobFamilys_url
  end

 
  
  test "should destroy jobFamily" do
    assert_difference("JobFamily.count", -1) do
      delete jobFamily_url(@jobFamily)
    end

    assert_redirected_to jobFamilys_url
  end
  
end


