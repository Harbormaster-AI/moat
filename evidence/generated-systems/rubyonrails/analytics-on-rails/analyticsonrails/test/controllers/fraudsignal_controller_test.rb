require "test_helper"

class FraudSignalControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @fraudSignal = fraudSignals(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create fraudSignal" do
    assert_difference("FraudSignal.count") do
      post fraudSignals_url, params: { fraudSignal: { name:"test string for name", ruleLogic:"test string for ruleLogic", SignalType:FraudSignal.SignalTypes[0] } }
    end

    assert_redirected_to fraudSignals_url
  end

 
  
  test "should destroy fraudSignal" do
    assert_difference("FraudSignal.count", -1) do
      delete fraudSignal_url(@fraudSignal)
    end

    assert_redirected_to fraudSignals_url
  end
  
end


