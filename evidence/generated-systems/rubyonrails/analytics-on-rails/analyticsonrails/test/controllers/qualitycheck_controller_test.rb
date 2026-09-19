require "test_helper"

class QualityCheckControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @qualityCheck = qualityChecks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create qualityCheck" do
    assert_difference("QualityCheck.count") do
      post qualityChecks_url, params: { qualityCheck: { checkedAt:1.week.ago, observedValue:"test value", sampleSize:100, Status:QualityCheck.Statuss[0] } }
    end

    assert_redirected_to qualityChecks_url
  end

 
  
  test "should destroy qualityCheck" do
    assert_difference("QualityCheck.count", -1) do
      delete qualityCheck_url(@qualityCheck)
    end

    assert_redirected_to qualityChecks_url
  end
  
end


