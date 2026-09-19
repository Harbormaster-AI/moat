require "test_helper"

class RoleControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @role = roles(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create role" do
    assert_difference("Role.count") do
      post roles_url, params: { role: { name:"test string for name", responsibility:"test string for responsibility" } }
    end

    assert_redirected_to roles_url
  end

 
  
  test "should destroy role" do
    assert_difference("Role.count", -1) do
      delete role_url(@role)
    end

    assert_redirected_to roles_url
  end
  
end


