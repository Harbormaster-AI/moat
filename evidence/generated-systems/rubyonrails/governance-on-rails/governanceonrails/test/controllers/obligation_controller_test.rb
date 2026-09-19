require "test_helper"

class ObligationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @obligation = obligations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create obligation" do
    assert_difference("Obligation.count") do
      post obligations_url, params: { obligation: { referenceNumber:"test string for referenceNumber", descriptionText:"test string for descriptionText", ObligationType:Obligation.ObligationTypes[0], ReviewFrequency:Obligation.ReviewFrequencys[0] } }
    end

    assert_redirected_to obligations_url
  end

 
  
  test "should destroy obligation" do
    assert_difference("Obligation.count", -1) do
      delete obligation_url(@obligation)
    end

    assert_redirected_to obligations_url
  end
  
end


