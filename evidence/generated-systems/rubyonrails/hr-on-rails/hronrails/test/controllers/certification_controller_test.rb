require "test_helper"

class CertificationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @certification = certifications(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create certification" do
    assert_difference("Certification.count") do
      post certifications_url, params: { certification: { name:"test string for name", issuer:"test string for issuer", validFrom:1.week.ago, validTo:1.week.ago, credentialId:"test string for credentialId" } }
    end

    assert_redirected_to certifications_url
  end

 
  
  test "should destroy certification" do
    assert_difference("Certification.count", -1) do
      delete certification_url(@certification)
    end

    assert_redirected_to certifications_url
  end
  
end


