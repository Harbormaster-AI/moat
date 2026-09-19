require "test_helper"

class ClaimPaymentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @claimPayment = claimPayments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create claimPayment" do
    assert_difference("ClaimPayment.count") do
      post claimPayments_url, params: { claimPayment: { paymentNumber:"test string for paymentNumber", amount:"test value", paymentDate:1.week.ago, PayeeType:ClaimPayment.PayeeTypes[0], Method:ClaimPayment.Methods[0], Status:ClaimPayment.Statuss[0] } }
    end

    assert_redirected_to claimPayments_url
  end

 
  
  test "should destroy claimPayment" do
    assert_difference("ClaimPayment.count", -1) do
      delete claimPayment_url(@claimPayment)
    end

    assert_redirected_to claimPayments_url
  end
  
end


