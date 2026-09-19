require "test_helper"

class ProductionCertificateControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productionCertificate = productionCertificates(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productionCertificate" do
    assert_difference("ProductionCertificate.count") do
      post productionCertificates_url, params: { productionCertificate: { certificateNumber:"test string for certificateNumber", authority:"test string for authority" } }
    end

    assert_redirected_to productionCertificates_url
  end

 
  
  test "should destroy productionCertificate" do
    assert_difference("ProductionCertificate.count", -1) do
      delete productionCertificate_url(@productionCertificate)
    end

    assert_redirected_to productionCertificates_url
  end
  
end


