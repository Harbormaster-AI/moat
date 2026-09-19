require "test_helper"

class InsurancePayerControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @insurancePayer = insurancePayers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create insurancePayer" do
    assert_difference("InsurancePayer.count") do
      post insurancePayers_url, params: { insurancePayer: { name:"test string for name", website:"test string for website", PayerType:InsurancePayer.PayerTypes[0] } }
    end

    assert_redirected_to insurancePayers_url
  end

 
  
  test "should destroy insurancePayer" do
    assert_difference("InsurancePayer.count", -1) do
      delete insurancePayer_url(@insurancePayer)
    end

    assert_redirected_to insurancePayers_url
  end
  
end


