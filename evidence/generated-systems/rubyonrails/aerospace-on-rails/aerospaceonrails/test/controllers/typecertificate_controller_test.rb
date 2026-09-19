require "test_helper"

class TypeCertificateControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @typeCertificate = typeCertificates(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create typeCertificate" do
    assert_difference("TypeCertificate.count") do
      post typeCertificates_url, params: { typeCertificate: { certificateNumber:"test string for certificateNumber", authority:"test string for authority" } }
    end

    assert_redirected_to typeCertificates_url
  end

 
  
  test "should destroy typeCertificate" do
    assert_difference("TypeCertificate.count", -1) do
      delete typeCertificate_url(@typeCertificate)
    end

    assert_redirected_to typeCertificates_url
  end
  
end


