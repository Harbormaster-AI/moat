require "test_helper"

class APUControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aPU = aPUs(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aPU" do
    assert_difference("APU.count") do
      post aPUs_url, params: { aPU: { model:"test string for model" } }
    end

    assert_redirected_to aPUs_url
  end

 
  
  test "should destroy aPU" do
    assert_difference("APU.count", -1) do
      delete aPU_url(@aPU)
    end

    assert_redirected_to aPUs_url
  end
  
end


