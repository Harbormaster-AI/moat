require "test_helper"

class PurchaseAgreementControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @purchaseAgreement = purchaseAgreements(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create purchaseAgreement" do
    assert_difference("PurchaseAgreement.count") do
      post purchaseAgreements_url, params: { purchaseAgreement: { agreementNumber:"test string for agreementNumber", effectiveDate:1.week.ago } }
    end

    assert_redirected_to purchaseAgreements_url
  end

 
  
  test "should destroy purchaseAgreement" do
    assert_difference("PurchaseAgreement.count", -1) do
      delete purchaseAgreement_url(@purchaseAgreement)
    end

    assert_redirected_to purchaseAgreements_url
  end
  
end


