require "test_helper"

class DocumentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @document = documents(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create document" do
    assert_difference("Document.count") do
      post documents_url, params: { document: { fileName:"test string for fileName", uploadedDate:1.week.ago, DocumentType:Document.DocumentTypes[0] } }
    end

    assert_redirected_to documents_url
  end

 
  
  test "should destroy document" do
    assert_difference("Document.count", -1) do
      delete document_url(@document)
    end

    assert_redirected_to documents_url
  end
  
end


