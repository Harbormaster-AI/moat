require "test_helper"

class InvestmentAccountControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @investmentAccount = investmentAccounts(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create investmentAccount" do
    assert_difference("InvestmentAccount.count") do
      post investmentAccounts_url, params: { investmentAccount: { accountNumber:"test value", baseCurrency:"test string for baseCurrency", balance:"test value", AccountType:InvestmentAccount.AccountTypes[0], Status:InvestmentAccount.Statuss[0] } }
    end

    assert_redirected_to investmentAccounts_url
  end

 
  
  test "should destroy investmentAccount" do
    assert_difference("InvestmentAccount.count", -1) do
      delete investmentAccount_url(@investmentAccount)
    end

    assert_redirected_to investmentAccounts_url
  end
  
end


