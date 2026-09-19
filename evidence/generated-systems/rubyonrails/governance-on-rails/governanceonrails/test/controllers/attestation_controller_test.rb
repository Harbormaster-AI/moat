require "test_helper"

class AttestationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @attestation = attestations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create attestation" do
    assert_difference("Attestation.count") do
      post attestations_url, params: { attestation: { statement:"test string for statement", attestor:"test string for attestor", dateSigned:1.week.ago, Result:Attestation.Results[0] } }
    end

    assert_redirected_to attestations_url
  end

 
  
  test "should destroy attestation" do
    assert_difference("Attestation.count", -1) do
      delete attestation_url(@attestation)
    end

    assert_redirected_to attestations_url
  end
  
end


