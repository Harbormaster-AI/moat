require "test_helper"

class PaymentContractControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @paymentContract = paymentContracts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create paymentContract" do
    assert_difference("PaymentContract.count") do
      post paymentContracts_url, params: { paymentContract: { contractNumber:"test string for contractNumber", pricingPlanCode:"test string for pricingPlanCode", Status:PaymentContract.Statuss[0] } }
    end

    assert_redirected_to paymentContracts_url
  end

 
  
  test "should destroy paymentContract" do
    assert_difference("PaymentContract.count", -1) do
      delete paymentContract_url(@paymentContract)
    end

    assert_redirected_to paymentContracts_url
  end
  
end


