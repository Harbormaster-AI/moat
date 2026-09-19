require "test_helper"

class ImagingReportControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @imagingReport = imagingReports(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create imagingReport" do
    assert_difference("ImagingReport.count") do
      post imagingReports_url, params: { imagingReport: { reportNumber:"test string for reportNumber", impression:"test string for impression", reportedDate:1.week.ago, Status:ImagingReport.Statuss[0] } }
    end

    assert_redirected_to imagingReports_url
  end

 
  
  test "should destroy imagingReport" do
    assert_difference("ImagingReport.count", -1) do
      delete imagingReport_url(@imagingReport)
    end

    assert_redirected_to imagingReports_url
  end
  
end


