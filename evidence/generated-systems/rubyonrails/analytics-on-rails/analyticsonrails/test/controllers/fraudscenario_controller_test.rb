require "test_helper"

class FraudScenarioControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @fraudScenario = fraudScenarios(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create fraudScenario" do
    assert_difference("FraudScenario.count") do
      post fraudScenarios_url, params: { fraudScenario: { name:"test string for name", riskAppetite:"test string for riskAppetite", DetectionType:FraudScenario.DetectionTypes[0] } }
    end

    assert_redirected_to fraudScenarios_url
  end

 
  
  test "should destroy fraudScenario" do
    assert_difference("FraudScenario.count", -1) do
      delete fraudScenario_url(@fraudScenario)
    end

    assert_redirected_to fraudScenarios_url
  end
  
end


