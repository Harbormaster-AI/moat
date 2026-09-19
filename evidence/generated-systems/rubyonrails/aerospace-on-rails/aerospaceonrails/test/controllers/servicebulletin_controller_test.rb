require "test_helper"

class ServiceBulletinControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @serviceBulletin = serviceBulletins(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create serviceBulletin" do
    assert_difference("ServiceBulletin.count") do
      post serviceBulletins_url, params: { serviceBulletin: { bulletinNumber:"test string for bulletinNumber", Category:ServiceBulletin.Categorys[0] } }
    end

    assert_redirected_to serviceBulletins_url
  end

 
  
  test "should destroy serviceBulletin" do
    assert_difference("ServiceBulletin.count", -1) do
      delete serviceBulletin_url(@serviceBulletin)
    end

    assert_redirected_to serviceBulletins_url
  end
  
end


