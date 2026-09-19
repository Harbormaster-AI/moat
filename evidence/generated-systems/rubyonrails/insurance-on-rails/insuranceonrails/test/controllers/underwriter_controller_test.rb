require "test_helper"

class UnderwriterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @underwriter = underwriters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create underwriter" do
    assert_difference("Underwriter.count") do
      post underwriters_url, params: { underwriter: { firstName:"test string for firstName", lastName:"test string for lastName", employeeId:"test string for employeeId", authorityLimit:"test value" } }
    end

    assert_redirected_to underwriters_url
  end

 
  
  test "should destroy underwriter" do
    assert_difference("Underwriter.count", -1) do
      delete underwriter_url(@underwriter)
    end

    assert_redirected_to underwriters_url
  end
  
end


