require "test_helper"

class LoanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @loan = loans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create loan" do
    assert_difference("Loan.count") do
      post loans_url, params: { loan: { loanNumber:"test string for loanNumber", principal:"test value", interestRate:"test value", originationDate:1.week.ago, maturityDate:1.week.ago, RateType:Loan.RateTypes[0], Status:Loan.Statuss[0] } }
    end

    assert_redirected_to loans_url
  end

 
  
  test "should destroy loan" do
    assert_difference("Loan.count", -1) do
      delete loan_url(@loan)
    end

    assert_redirected_to loans_url
  end
  
end


