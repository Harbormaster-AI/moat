require "test_helper"

class UoMConversionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @uoMConversion = uoMConversions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create uoMConversion" do
    assert_difference("UoMConversion.count") do
      post uoMConversions_url, params: { uoMConversion: { factor:"test value", precision:100, FromUnit:UoMConversion.FromUnits[0], ToUnit:UoMConversion.ToUnits[0] } }
    end

    assert_redirected_to uoMConversions_url
  end

 
  
  test "should destroy uoMConversion" do
    assert_difference("UoMConversion.count", -1) do
      delete uoMConversion_url(@uoMConversion)
    end

    assert_redirected_to uoMConversions_url
  end
  
end


