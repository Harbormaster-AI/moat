require "test_helper"

class RegulationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @regulation = regulations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create regulation" do
    assert_difference("Regulation.count") do
      post regulations_url, params: { regulation: { name:"test string for name", citation:"test string for citation", jurisdiction:"test string for jurisdiction", publicationUrl:"test value" } }
    end

    assert_redirected_to regulations_url
  end

 
  
  test "should destroy regulation" do
    assert_difference("Regulation.count", -1) do
      delete regulation_url(@regulation)
    end

    assert_redirected_to regulations_url
  end
  
end


