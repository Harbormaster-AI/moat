require "test_helper"

class InsuredObjectControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @insuredObject = insuredObjects(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create insuredObject" do
    assert_difference("InsuredObject.count") do
      post insuredObjects_url, params: { insuredObject: { description:"test string for description", serialOrId:"test string for serialOrId", primaryAddress:"test value", ObjectType:InsuredObject.ObjectTypes[0] } }
    end

    assert_redirected_to insuredObjects_url
  end

 
  
  test "should destroy insuredObject" do
    assert_difference("InsuredObject.count", -1) do
      delete insuredObject_url(@insuredObject)
    end

    assert_redirected_to insuredObjects_url
  end
  
end


