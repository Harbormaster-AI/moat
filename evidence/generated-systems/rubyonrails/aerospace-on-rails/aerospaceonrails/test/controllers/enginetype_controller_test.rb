require "test_helper"

class EngineTypeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @engineType = engineTypes(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create engineType" do
    assert_difference("EngineType.count") do
      post engineTypes_url, params: { engineType: { engineModelCode:"test string for engineModelCode", maxThrustKn:"test value", Category:EngineType.Categorys[0] } }
    end

    assert_redirected_to engineTypes_url
  end

 
  
  test "should destroy engineType" do
    assert_difference("EngineType.count", -1) do
      delete engineType_url(@engineType)
    end

    assert_redirected_to engineTypes_url
  end
  
end


