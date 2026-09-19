require "test_helper"

class OperatorControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @operator = operators(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create operator" do
    assert_difference("Operator.count") do
      post operators_url, params: { operator: { name:"test string for name", icaoDesignator:"test string for icaoDesignator", OperatorType:Operator.OperatorTypes[0] } }
    end

    assert_redirected_to operators_url
  end

 
  
  test "should destroy operator" do
    assert_difference("Operator.count", -1) do
      delete operator_url(@operator)
    end

    assert_redirected_to operators_url
  end
  
end


