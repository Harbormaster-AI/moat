require "test_helper"

class DataCategoryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataCategory = dataCategorys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataCategory" do
    assert_difference("DataCategory.count") do
      post dataCategorys_url, params: { dataCategory: { name:"test string for name", description:"test string for description", Classification:DataCategory.Classifications[0] } }
    end

    assert_redirected_to dataCategorys_url
  end

 
  
  test "should destroy dataCategory" do
    assert_difference("DataCategory.count", -1) do
      delete dataCategory_url(@dataCategory)
    end

    assert_redirected_to dataCategorys_url
  end
  
end


