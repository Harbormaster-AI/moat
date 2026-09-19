require "test_helper"

class ReinsuranceAgreementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @reinsuranceAgreement = reinsuranceAgreements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create reinsuranceAgreement" do
    assert_difference("ReinsuranceAgreement.count") do
      post reinsuranceAgreements_url, params: { reinsuranceAgreement: { agreementNumber:"test string for agreementNumber", effectivePeriod:1.week.ago, retention:"test value", limit:"test value", cessionPercentage:"test value", ReinsuranceType:ReinsuranceAgreement.ReinsuranceTypes[0], TreatyType:ReinsuranceAgreement.TreatyTypes[0] } }
    end

    assert_redirected_to reinsuranceAgreements_url
  end

 
  
  test "should destroy reinsuranceAgreement" do
    assert_difference("ReinsuranceAgreement.count", -1) do
      delete reinsuranceAgreement_url(@reinsuranceAgreement)
    end

    assert_redirected_to reinsuranceAgreements_url
  end
  
end


