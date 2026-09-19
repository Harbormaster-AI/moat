require "test_helper"

class AdmissionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @admission = admissions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create admission" do
    assert_difference("Admission.count") do
      post admissions_url, params: { admission: { admitDateTime:1.week.ago, bed:"test string for bed", AdmissionType:Admission.AdmissionTypes[0] } }
    end

    assert_redirected_to admissions_url
  end

 
  
  test "should destroy admission" do
    assert_difference("Admission.count", -1) do
      delete admission_url(@admission)
    end

    assert_redirected_to admissions_url
  end
  
end


