require "test_helper"

class PerformanceReviewControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @performanceReview = performanceReviews(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create performanceReview" do
    assert_difference("PerformanceReview.count") do
      post performanceReviews_url, params: { performanceReview: { reviewNumber:"test string for reviewNumber", reviewDate:1.week.ago, reviewerComments:"test string for reviewerComments", Rating:PerformanceReview.Ratings[0], Status:PerformanceReview.Statuss[0] } }
    end

    assert_redirected_to performanceReviews_url
  end

 
  
  test "should destroy performanceReview" do
    assert_difference("PerformanceReview.count", -1) do
      delete performanceReview_url(@performanceReview)
    end

    assert_redirected_to performanceReviews_url
  end
  
end


