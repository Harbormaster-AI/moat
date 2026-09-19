require "test_helper"

class OfferControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @offer = offers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create offer" do
    assert_difference("Offer.count") do
      post offers_url, params: { offer: { offerNumber:"test string for offerNumber", proposedStartDate:1.week.ago, baseSalary:"test value", signOnBonus:"test value", Status:Offer.Statuss[0] } }
    end

    assert_redirected_to offers_url
  end

 
  
  test "should destroy offer" do
    assert_difference("Offer.count", -1) do
      delete offer_url(@offer)
    end

    assert_redirected_to offers_url
  end
  
end


