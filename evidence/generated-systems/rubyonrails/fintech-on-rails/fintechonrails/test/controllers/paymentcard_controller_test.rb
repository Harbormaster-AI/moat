require "test_helper"

class PaymentCardControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @paymentCard = paymentCards(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create paymentCard" do
    assert_difference("PaymentCard.count") do
      post paymentCards_url, params: { paymentCard: { cardToken:"test value", maskedPan:"test string for maskedPan", expiryMonth:100, expiryYear:100, cardholderName:"test string for cardholderName", Scheme:PaymentCard.Schemes[0], Status:PaymentCard.Statuss[0] } }
    end

    assert_redirected_to paymentCards_url
  end

 
  
  test "should destroy paymentCard" do
    assert_difference("PaymentCard.count", -1) do
      delete paymentCard_url(@paymentCard)
    end

    assert_redirected_to paymentCards_url
  end
  
end


