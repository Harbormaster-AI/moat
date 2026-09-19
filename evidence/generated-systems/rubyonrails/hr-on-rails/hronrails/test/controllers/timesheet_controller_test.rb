require "test_helper"

class TimesheetControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @timesheet = timesheets(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create timesheet" do
    assert_difference("Timesheet.count") do
      post timesheets_url, params: { timesheet: { periodStart:1.week.ago, periodEnd:1.week.ago, submissionDate:1.week.ago, Status:Timesheet.Statuss[0] } }
    end

    assert_redirected_to timesheets_url
  end

 
  
  test "should destroy timesheet" do
    assert_difference("Timesheet.count", -1) do
      delete timesheet_url(@timesheet)
    end

    assert_redirected_to timesheets_url
  end
  
end


