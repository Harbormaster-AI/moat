require "test_helper"

class OperationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @operation = operations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create operation" do
    assert_difference("Operation.count") do
      post operations_url, params: { operation: { operationNumber:"test string for operationNumber", name:"test string for name", setupTime:1.week.ago, standardCycleTime:1.week.ago, OperationType:Operation.OperationTypes[0] } }
    end

    assert_redirected_to operations_url
  end

 
  
  test "should destroy operation" do
    assert_difference("Operation.count", -1) do
      delete operation_url(@operation)
    end

    assert_redirected_to operations_url
  end
  
end


