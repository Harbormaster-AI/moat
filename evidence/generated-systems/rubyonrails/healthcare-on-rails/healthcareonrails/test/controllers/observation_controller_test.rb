require "test_helper"

class ObservationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @observation = observations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create observation" do
    assert_difference("Observation.count") do
      post observations_url, params: { observation: { code:"test string for code", value:"test string for value", unit:"test string for unit", effectiveDateTime:1.week.ago, Interpretation:Observation.Interpretations[0] } }
    end

    assert_redirected_to observations_url
  end

 
  
  test "should destroy observation" do
    assert_difference("Observation.count", -1) do
      delete observation_url(@observation)
    end

    assert_redirected_to observations_url
  end
  
end


