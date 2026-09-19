require "test_helper"

class InferenceEndpointControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @inferenceEndpoint = inferenceEndpoints(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create inferenceEndpoint" do
    assert_difference("InferenceEndpoint.count") do
      post inferenceEndpoints_url, params: { inferenceEndpoint: { name:"test string for name", endpointUrl:"test string for endpointUrl", trafficShare:"test value", Mode:InferenceEndpoint.Modes[0] } }
    end

    assert_redirected_to inferenceEndpoints_url
  end

 
  
  test "should destroy inferenceEndpoint" do
    assert_difference("InferenceEndpoint.count", -1) do
      delete inferenceEndpoint_url(@inferenceEndpoint)
    end

    assert_redirected_to inferenceEndpoints_url
  end
  
end


