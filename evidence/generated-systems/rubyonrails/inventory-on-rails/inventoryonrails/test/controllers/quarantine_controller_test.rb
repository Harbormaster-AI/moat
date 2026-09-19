require "test_helper"

class QuarantineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @quarantine = quarantines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create quarantine" do
    assert_difference("Quarantine.count") do
      post quarantines_url, params: { quarantine: { reason:"test string for reason", startedAt:1.week.ago, releasedAt:1.week.ago, Disposition:Quarantine.Dispositions[0] } }
    end

    assert_redirected_to quarantines_url
  end

 
  
  test "should destroy quarantine" do
    assert_difference("Quarantine.count", -1) do
      delete quarantine_url(@quarantine)
    end

    assert_redirected_to quarantines_url
  end
  
end


