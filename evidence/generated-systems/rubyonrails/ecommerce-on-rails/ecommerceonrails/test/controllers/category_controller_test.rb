require "test_helper"

class CategoryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @category = categorys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create category" do
    assert_difference("Category.count") do
      post categorys_url, params: { category: { name:"test string for name", slug:"test string for slug", position:100, asActive:true } }
    end

    assert_redirected_to categorys_url
  end

 
  
  test "should destroy category" do
    assert_difference("Category.count", -1) do
      delete category_url(@category)
    end

    assert_redirected_to categorys_url
  end
  
end


