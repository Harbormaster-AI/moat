require "test_helper"

class ThirdPartyAssessmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @thirdPartyAssessment = thirdPartyAssessments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create thirdPartyAssessment" do
    assert_difference("ThirdPartyAssessment.count") do
      post thirdPartyAssessments_url, params: { thirdPartyAssessment: { assessmentDate:1.week.ago, assessor:"test string for assessor", AssessmentType:ThirdPartyAssessment.AssessmentTypes[0], Result:ThirdPartyAssessment.Results[0] } }
    end

    assert_redirected_to thirdPartyAssessments_url
  end

 
  
  test "should destroy thirdPartyAssessment" do
    assert_difference("ThirdPartyAssessment.count", -1) do
      delete thirdPartyAssessment_url(@thirdPartyAssessment)
    end

    assert_redirected_to thirdPartyAssessments_url
  end
  
end


