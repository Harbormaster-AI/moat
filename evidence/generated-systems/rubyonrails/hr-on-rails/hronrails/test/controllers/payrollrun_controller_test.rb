require "test_helper"

class PayrollRunControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @payrollRun = payrollRuns(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create payrollRun" do
    assert_difference("PayrollRun.count") do
      post payrollRuns_url, params: { payrollRun: { runNumber:"test string for runNumber", periodStart:1.week.ago, periodEnd:1.week.ago, paymentDate:1.week.ago, Status:PayrollRun.Statuss[0] } }
    end

    assert_redirected_to payrollRuns_url
  end

 
  
  test "should destroy payrollRun" do
    assert_difference("PayrollRun.count", -1) do
      delete payrollRun_url(@payrollRun)
    end

    assert_redirected_to payrollRuns_url
  end
  
end


