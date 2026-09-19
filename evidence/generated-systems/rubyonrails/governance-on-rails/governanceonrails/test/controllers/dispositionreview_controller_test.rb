require "test_helper"

class DispositionReviewControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dispositionReview = dispositionReviews(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dispositionReview" do
    assert_difference("DispositionReview.count") do
      post dispositionReviews_url, params: { dispositionReview: { reviewDate:1.week.ago, reviewer:"test string for reviewer", notes:"test string for notes", Outcome:DispositionReview.Outcomes[0] } }
    end

    assert_redirected_to dispositionReviews_url
  end

 
  
  test "should destroy dispositionReview" do
    assert_difference("DispositionReview.count", -1) do
      delete dispositionReview_url(@dispositionReview)
    end

    assert_redirected_to dispositionReviews_url
  end
  
end


