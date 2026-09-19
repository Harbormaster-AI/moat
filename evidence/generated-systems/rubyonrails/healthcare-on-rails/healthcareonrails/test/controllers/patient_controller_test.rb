require "test_helper"

class PatientControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @patient = patients(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create patient" do
    assert_difference("Patient.count") do
      post patients_url, params: { patient: { firstName:"test string for firstName", lastName:"test string for lastName", mrn:"test value", dateOfBirth:1.week.ago, address:"test value", primaryLanguage:"test string for primaryLanguage", SexAtBirth:Patient.SexAtBirths[0], BloodType:Patient.BloodTypes[0] } }
    end

    assert_redirected_to patients_url
  end

 
  
  test "should destroy patient" do
    assert_difference("Patient.count", -1) do
      delete patient_url(@patient)
    end

    assert_redirected_to patients_url
  end
  
end


