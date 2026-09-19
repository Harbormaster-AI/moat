require "test_helper"

class LegalHoldControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @legalHold = legalHolds(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create legalHold" do
    assert_difference("LegalHold.count") do
      post legalHolds_url, params: { legalHold: { name:"test string for name", reason:"test string for reason", issuedDate:1.week.ago, releaseDate:1.week.ago, HoldStatus:LegalHold.HoldStatuss[0] } }
    end

    assert_redirected_to legalHolds_url
  end

 
  
  test "should destroy legalHold" do
    assert_difference("LegalHold.count", -1) do
      delete legalHold_url(@legalHold)
    end

    assert_redirected_to legalHolds_url
  end
  
end


