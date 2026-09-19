require "test_helper"

class BankAccountControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @bankAccount = bankAccounts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create bankAccount" do
    assert_difference("BankAccount.count") do
      post bankAccounts_url, params: { bankAccount: { accountHolder:"test string for accountHolder", bankName:"test string for bankName", iban:"test string for iban", bic:"test string for bic", accountNumber:"test string for accountNumber", routingNumber:"test string for routingNumber" } }
    end

    assert_redirected_to bankAccounts_url
  end

 
  
  test "should destroy bankAccount" do
    assert_difference("BankAccount.count", -1) do
      delete bankAccount_url(@bankAccount)
    end

    assert_redirected_to bankAccounts_url
  end
  
end


