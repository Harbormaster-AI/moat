require "test_helper"

class LoanTransactionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @loanTransaction = loanTransactions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create loanTransaction" do
    assert_difference("LoanTransaction.count") do
      post loanTransactions_url, params: { loanTransaction: { transactionId:"test value", amount:"test value", postingDate:1.week.ago, Type:LoanTransaction.Types[0], Status:LoanTransaction.Statuss[0] } }
    end

    assert_redirected_to loanTransactions_url
  end

 
  
  test "should destroy loanTransaction" do
    assert_difference("LoanTransaction.count", -1) do
      delete loanTransaction_url(@loanTransaction)
    end

    assert_redirected_to loanTransactions_url
  end
  
end


