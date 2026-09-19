require "test_helper"

class EvidenceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @evidence = evidences(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create evidence" do
    assert_difference("Evidence.count") do
      post evidences_url, params: { evidence: { title:"test string for title", locationUrl:"test value", receivedDate:1.week.ago, EvidenceType:Evidence.EvidenceTypes[0] } }
    end

    assert_redirected_to evidences_url
  end

 
  
  test "should destroy evidence" do
    assert_difference("Evidence.count", -1) do
      delete evidence_url(@evidence)
    end

    assert_redirected_to evidences_url
  end
  
end


