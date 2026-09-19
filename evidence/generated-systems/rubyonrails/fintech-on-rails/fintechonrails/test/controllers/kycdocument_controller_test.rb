require "test_helper"

class KYCDocumentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @kYCDocument = kYCDocuments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create kYCDocument" do
    assert_difference("KYCDocument.count") do
      post kYCDocuments_url, params: { kYCDocument: { reference:"test value", issuedCountry:"test string for issuedCountry", expirationDate:1.week.ago, DocumentType:KYCDocument.DocumentTypes[0], Status:KYCDocument.Statuss[0] } }
    end

    assert_redirected_to kYCDocuments_url
  end

 
  
  test "should destroy kYCDocument" do
    assert_difference("KYCDocument.count", -1) do
      delete kYCDocument_url(@kYCDocument)
    end

    assert_redirected_to kYCDocuments_url
  end
  
end


