require "test_helper"

class BeneficiaryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @beneficiary = beneficiarys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create beneficiary" do
    assert_difference("Beneficiary.count") do
      post beneficiarys_url, params: { beneficiary: { name:"test string for name", share:"test value", Relationship:Beneficiary.Relationships[0] } }
    end

    assert_redirected_to beneficiarys_url
  end

 
  
  test "should destroy beneficiary" do
    assert_difference("Beneficiary.count", -1) do
      delete beneficiary_url(@beneficiary)
    end

    assert_redirected_to beneficiarys_url
  end
  
end


