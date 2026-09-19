require "test_helper"

class PayrollCalendarControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @payrollCalendar = payrollCalendars(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create payrollCalendar" do
    assert_difference("PayrollCalendar.count") do
      post payrollCalendars_url, params: { payrollCalendar: { name:"test string for name", country:"test string for country", PayFrequency:PayrollCalendar.PayFrequencys[0] } }
    end

    assert_redirected_to payrollCalendars_url
  end

 
  
  test "should destroy payrollCalendar" do
    assert_difference("PayrollCalendar.count", -1) do
      delete payrollCalendar_url(@payrollCalendar)
    end

    assert_redirected_to payrollCalendars_url
  end
  
end


