require "test_helper"

class AgreementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @agreement = agreements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create agreement" do
    assert_difference("Agreement.count") do
      post agreements_url, params: { agreement: { agreementNumber:"test string for agreementNumber", effectiveDate:1.week.ago, AgreementType:Agreement.AgreementTypes[0], Status:Agreement.Statuss[0] } }
    end

    assert_redirected_to agreements_url
  end

 
  
  test "should destroy agreement" do
    assert_difference("Agreement.count", -1) do
      delete agreement_url(@agreement)
    end

    assert_redirected_to agreements_url
  end
  
end


