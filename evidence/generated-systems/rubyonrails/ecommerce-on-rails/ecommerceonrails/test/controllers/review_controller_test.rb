require "test_helper"

class ReviewControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @review = reviews(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create review" do
    assert_difference("Review.count") do
      post reviews_url, params: { review: { rating:100, title:"test string for title", content:"test string for content", createdAt:1.week.ago, Status:Review.Statuss[0] } }
    end

    assert_redirected_to reviews_url
  end

 
  
  test "should destroy review" do
    assert_difference("Review.count", -1) do
      delete review_url(@review)
    end

    assert_redirected_to reviews_url
  end
  
end


