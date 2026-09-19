require "test_helper"

class PolicyAcknowledgementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @policyAcknowledgement = policyAcknowledgements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create policyAcknowledgement" do
    assert_difference("PolicyAcknowledgement.count") do
      post policyAcknowledgements_url, params: { policyAcknowledgement: { acknowledgementDate:1.week.ago, Status:PolicyAcknowledgement.Statuss[0] } }
    end

    assert_redirected_to policyAcknowledgements_url
  end

 
  
  test "should destroy policyAcknowledgement" do
    assert_difference("PolicyAcknowledgement.count", -1) do
      delete policyAcknowledgement_url(@policyAcknowledgement)
    end

    assert_redirected_to policyAcknowledgements_url
  end
  
end


