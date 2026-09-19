require "test_helper"

class DiagnosisControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @diagnosis = diagnosiss(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create diagnosis" do
    assert_difference("Diagnosis.count") do
      post diagnosiss_url, params: { diagnosis: { code:"test string for code", description:"test string for description", onsetDate:1.week.ago, Certainty:Diagnosis.Certaintys[0] } }
    end

    assert_redirected_to diagnosiss_url
  end

 
  
  test "should destroy diagnosis" do
    assert_difference("Diagnosis.count", -1) do
      delete diagnosis_url(@diagnosis)
    end

    assert_redirected_to diagnosiss_url
  end
  
end


