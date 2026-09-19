require "test_helper"

class PositionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @position = positions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create position" do
    assert_difference("Position.count") do
      post positions_url, params: { position: { quantity:"test value", averageCost:"test value", marketValue:"test value" } }
    end

    assert_redirected_to positions_url
  end

 
  
  test "should destroy position" do
    assert_difference("Position.count", -1) do
      delete position_url(@position)
    end

    assert_redirected_to positions_url
  end
  
end


