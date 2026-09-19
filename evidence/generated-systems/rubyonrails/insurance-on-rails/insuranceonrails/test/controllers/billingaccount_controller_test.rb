require "test_helper"

class BillingAccountControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @billingAccount = billingAccounts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create billingAccount" do
    assert_difference("BillingAccount.count") do
      post billingAccounts_url, params: { billingAccount: { accountNumber:"test string for accountNumber", balance:"test value", Status:BillingAccount.Statuss[0] } }
    end

    assert_redirected_to billingAccounts_url
  end

 
  
  test "should destroy billingAccount" do
    assert_difference("BillingAccount.count", -1) do
      delete billingAccount_url(@billingAccount)
    end

    assert_redirected_to billingAccounts_url
  end
  
end


