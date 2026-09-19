require "test_helper"

class EndorsementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @endorsement = endorsements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create endorsement" do
    assert_difference("Endorsement.count") do
      post endorsements_url, params: { endorsement: { endorsementNumber:"test string for endorsementNumber", effectiveDate:1.week.ago, description:"test string for description" } }
    end

    assert_redirected_to endorsements_url
  end

 
  
  test "should destroy endorsement" do
    assert_difference("Endorsement.count", -1) do
      delete endorsement_url(@endorsement)
    end

    assert_redirected_to endorsements_url
  end
  
end


