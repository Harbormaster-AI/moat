require "test_helper"

class AvionicsSuiteControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @avionicsSuite = avionicsSuites(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create avionicsSuite" do
    assert_difference("AvionicsSuite.count") do
      post avionicsSuites_url, params: { avionicsSuite: { suiteName:"test string for suiteName", softwareBaseline:"test string for softwareBaseline" } }
    end

    assert_redirected_to avionicsSuites_url
  end

 
  
  test "should destroy avionicsSuite" do
    assert_difference("AvionicsSuite.count", -1) do
      delete avionicsSuite_url(@avionicsSuite)
    end

    assert_redirected_to avionicsSuites_url
  end
  
end


