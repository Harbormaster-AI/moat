require "test_helper"

class TerminalControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @terminal = terminals(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create terminal" do
    assert_difference("Terminal.count") do
      post terminals_url, params: { terminal: { location:"test value", Type:Terminal.Types[0], Status:Terminal.Statuss[0] } }
    end

    assert_redirected_to terminals_url
  end

 
  
  test "should destroy terminal" do
    assert_difference("Terminal.count", -1) do
      delete terminal_url(@terminal)
    end

    assert_redirected_to terminals_url
  end
  
end


