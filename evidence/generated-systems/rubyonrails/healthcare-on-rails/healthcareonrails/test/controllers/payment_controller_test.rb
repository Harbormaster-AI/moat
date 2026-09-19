require "test_helper"

class PaymentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @payment = payments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create payment" do
    assert_difference("Payment.count") do
      post payments_url, params: { payment: { paymentNumber:"test string for paymentNumber", amount:"test value", paymentDate:1.week.ago, Method:Payment.Methods[0] } }
    end

    assert_redirected_to payments_url
  end

 
  
  test "should destroy payment" do
    assert_difference("Payment.count", -1) do
      delete payment_url(@payment)
    end

    assert_redirected_to payments_url
  end
  
end


