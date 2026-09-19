require "test_helper"

class BIQueryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @bIQuery = bIQuerys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create bIQuery" do
    assert_difference("BIQuery.count") do
      post bIQuerys_url, params: { bIQuery: { name:"test string for name", text:"test string for text", Dialect:BIQuery.Dialects[0] } }
    end

    assert_redirected_to bIQuerys_url
  end

 
  
  test "should destroy bIQuery" do
    assert_difference("BIQuery.count", -1) do
      delete bIQuery_url(@bIQuery)
    end

    assert_redirected_to bIQuerys_url
  end
  
end


