require "test_helper"

class CatalogControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @catalog = catalogs(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create catalog" do
    assert_difference("Catalog.count") do
      post catalogs_url, params: { catalog: { name:"test string for name", catalogCode:"test string for catalogCode", asActive:true } }
    end

    assert_redirected_to catalogs_url
  end

 
  
  test "should destroy catalog" do
    assert_difference("Catalog.count", -1) do
      delete catalog_url(@catalog)
    end

    assert_redirected_to catalogs_url
  end
  
end


