require "test_helper"

class MatterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @matter = matters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create matter" do
    assert_difference("Matter.count") do
      post matters_url, params: { matter: { matterName:"test string for matterName", leadCounsel:"test string for leadCounsel", MatterType:Matter.MatterTypes[0], Status:Matter.Statuss[0] } }
    end

    assert_redirected_to matters_url
  end

 
  
  test "should destroy matter" do
    assert_difference("Matter.count", -1) do
      delete matter_url(@matter)
    end

    assert_redirected_to matters_url
  end
  
end


