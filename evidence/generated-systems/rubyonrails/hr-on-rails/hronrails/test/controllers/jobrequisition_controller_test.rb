require "test_helper"

class JobRequisitionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @jobRequisition = jobRequisitions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create jobRequisition" do
    assert_difference("JobRequisition.count") do
      post jobRequisitions_url, params: { jobRequisition: { requisitionNumber:"test string for requisitionNumber", title:"test string for title", openings:100, targetStartDate:1.week.ago, Status:JobRequisition.Statuss[0], Priority:JobRequisition.Prioritys[0] } }
    end

    assert_redirected_to jobRequisitions_url
  end

 
  
  test "should destroy jobRequisition" do
    assert_difference("JobRequisition.count", -1) do
      delete jobRequisition_url(@jobRequisition)
    end

    assert_redirected_to jobRequisitions_url
  end
  
end


