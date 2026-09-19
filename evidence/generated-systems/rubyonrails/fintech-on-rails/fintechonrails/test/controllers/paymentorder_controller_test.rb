require "test_helper"

class PaymentOrderControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @paymentOrder = paymentOrders(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create paymentOrder" do
    assert_difference("PaymentOrder.count") do
      post paymentOrders_url, params: { paymentOrder: { orderReference:"test string for orderReference", requestedExecutionDate:1.week.ago, purpose:"test string for purpose", PaymentMethod:PaymentOrder.PaymentMethods[0], Status:PaymentOrder.Statuss[0], Priority:PaymentOrder.Prioritys[0] } }
    end

    assert_redirected_to paymentOrders_url
  end

 
  
  test "should destroy paymentOrder" do
    assert_difference("PaymentOrder.count", -1) do
      delete paymentOrder_url(@paymentOrder)
    end

    assert_redirected_to paymentOrders_url
  end
  
end


