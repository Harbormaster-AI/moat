require "test_helper"

class BenefitEnrollmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @benefitEnrollment = benefitEnrollments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create benefitEnrollment" do
    assert_difference("BenefitEnrollment.count") do
      post benefitEnrollments_url, params: { benefitEnrollment: { enrollmentId:"test string for enrollmentId", effectiveFrom:1.week.ago, effectiveTo:1.week.ago, Status:BenefitEnrollment.Statuss[0], CoverageLevel:BenefitEnrollment.CoverageLevels[0] } }
    end

    assert_redirected_to benefitEnrollments_url
  end

 
  
  test "should destroy benefitEnrollment" do
    assert_difference("BenefitEnrollment.count", -1) do
      delete benefitEnrollment_url(@benefitEnrollment)
    end

    assert_redirected_to benefitEnrollments_url
  end
  
end


