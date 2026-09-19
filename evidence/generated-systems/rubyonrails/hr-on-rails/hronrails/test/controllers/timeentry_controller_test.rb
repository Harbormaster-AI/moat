require "test_helper"

class TimeEntryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @timeEntry = timeEntrys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create timeEntry" do
    assert_difference("TimeEntry.count") do
      post timeEntrys_url, params: { timeEntry: { entryDate:1.week.ago, hoursWorked:"test value", EntryType:TimeEntry.EntryTypes[0] } }
    end

    assert_redirected_to timeEntrys_url
  end

 
  
  test "should destroy timeEntry" do
    assert_difference("TimeEntry.count", -1) do
      delete timeEntry_url(@timeEntry)
    end

    assert_redirected_to timeEntrys_url
  end
  
end


