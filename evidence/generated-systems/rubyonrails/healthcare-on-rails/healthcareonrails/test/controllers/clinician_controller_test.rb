require "test_helper"

class ClinicianControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @clinician = clinicians(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create clinician" do
    assert_difference("Clinician.count") do
      post clinicians_url, params: { clinician: { firstName:"test string for firstName", lastName:"test string for lastName", licenseNumber:"test string for licenseNumber", ClinicianType:Clinician.ClinicianTypes[0], Specialty:Clinician.Specialtys[0] } }
    end

    assert_redirected_to clinicians_url
  end

 
  
  test "should destroy clinician" do
    assert_difference("Clinician.count", -1) do
      delete clinician_url(@clinician)
    end

    assert_redirected_to clinicians_url
  end
  
end


