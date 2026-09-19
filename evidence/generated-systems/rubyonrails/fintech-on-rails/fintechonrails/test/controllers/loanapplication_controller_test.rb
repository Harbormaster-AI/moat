require "test_helper"

class LoanApplicationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @loanApplication = loanApplications(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create loanApplication" do
    assert_difference("LoanApplication.count") do
      post loanApplications_url, params: { loanApplication: { applicationNumber:"test string for applicationNumber", amountRequested:"test value", termMonths:100, submittedAt:1.week.ago, Product:LoanApplication.Products[0], Purpose:LoanApplication.Purposes[0], Status:LoanApplication.Statuss[0] } }
    end

    assert_redirected_to loanApplications_url
  end

 
  
  test "should destroy loanApplication" do
    assert_difference("LoanApplication.count", -1) do
      delete loanApplication_url(@loanApplication)
    end

    assert_redirected_to loanApplications_url
  end
  
end


