require "test_helper"

class PrivacyNoticeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @privacyNotice = privacyNotices(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create privacyNotice" do
    assert_difference("PrivacyNotice.count") do
      post privacyNotices_url, params: { privacyNotice: { title:"test string for title", audience:"test string for audience", versionLabel:"test string for versionLabel", publicationDate:1.week.ago, publicationUrl:"test value", Status:PrivacyNotice.Statuss[0] } }
    end

    assert_redirected_to privacyNotices_url
  end

 
  
  test "should destroy privacyNotice" do
    assert_difference("PrivacyNotice.count", -1) do
      delete privacyNotice_url(@privacyNotice)
    end

    assert_redirected_to privacyNotices_url
  end
  
end


