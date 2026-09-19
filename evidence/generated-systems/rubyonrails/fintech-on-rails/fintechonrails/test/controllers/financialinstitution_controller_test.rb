require "test_helper"

class FinancialInstitutionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @financialInstitution = financialInstitutions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create financialInstitution" do
    assert_difference("FinancialInstitution.count") do
      post financialInstitutions_url, params: { financialInstitution: { name:"test string for name", legalName:"test string for legalName", countryOfIncorporation:"test string for countryOfIncorporation", bic:"test value", website:"test string for website" } }
    end

    assert_redirected_to financialInstitutions_url
  end

 
  
  test "should destroy financialInstitution" do
    assert_difference("FinancialInstitution.count", -1) do
      delete financialInstitution_url(@financialInstitution)
    end

    assert_redirected_to financialInstitutions_url
  end
  
end


