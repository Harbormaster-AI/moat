require "test_helper"

class PaymentProcessorControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @paymentProcessor = paymentProcessors(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create paymentProcessor" do
    assert_difference("PaymentProcessor.count") do
      post paymentProcessors_url, params: { paymentProcessor: { name:"test string for name", processorCode:"test string for processorCode", networkSupport:"test string for networkSupport" } }
    end

    assert_redirected_to paymentProcessors_url
  end

 
  
  test "should destroy paymentProcessor" do
    assert_difference("PaymentProcessor.count", -1) do
      delete paymentProcessor_url(@paymentProcessor)
    end

    assert_redirected_to paymentProcessors_url
  end
  
end


