require "test_helper"

class DirectDebitMandateControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @directDebitMandate = directDebitMandates(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create directDebitMandate" do
    assert_difference("DirectDebitMandate.count") do
      post directDebitMandates_url, params: { directDebitMandate: { mandateId:"test string for mandateId", signedAt:1.week.ago, Scheme:DirectDebitMandate.Schemes[0], Status:DirectDebitMandate.Statuss[0] } }
    end

    assert_redirected_to directDebitMandates_url
  end

 
  
  test "should destroy directDebitMandate" do
    assert_difference("DirectDebitMandate.count", -1) do
      delete directDebitMandate_url(@directDebitMandate)
    end

    assert_redirected_to directDebitMandates_url
  end
  
end


