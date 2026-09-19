require "test_helper"

class AuditWorkpaperControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @auditWorkpaper = auditWorkpapers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create auditWorkpaper" do
    assert_difference("AuditWorkpaper.count") do
      post auditWorkpapers_url, params: { auditWorkpaper: { workpaperRef:"test string for workpaperRef", subject:"test string for subject", workpaperUrl:"test value" } }
    end

    assert_redirected_to auditWorkpapers_url
  end

 
  
  test "should destroy auditWorkpaper" do
    assert_difference("AuditWorkpaper.count", -1) do
      delete auditWorkpaper_url(@auditWorkpaper)
    end

    assert_redirected_to auditWorkpapers_url
  end
  
end


