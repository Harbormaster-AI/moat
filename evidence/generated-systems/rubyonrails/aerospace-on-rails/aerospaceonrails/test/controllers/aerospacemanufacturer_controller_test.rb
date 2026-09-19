require "test_helper"

class AerospaceManufacturerControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @aerospaceManufacturer = aerospaceManufacturers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create aerospaceManufacturer" do
    assert_difference("AerospaceManufacturer.count") do
      post aerospaceManufacturers_url, params: { aerospaceManufacturer: { name:"test string for name", legalName:"test string for legalName", headquartersCountry:"test string for headquartersCountry", website:"test string for website" } }
    end

    assert_redirected_to aerospaceManufacturers_url
  end

 
  
  test "should destroy aerospaceManufacturer" do
    assert_difference("AerospaceManufacturer.count", -1) do
      delete aerospaceManufacturer_url(@aerospaceManufacturer)
    end

    assert_redirected_to aerospaceManufacturers_url
  end
  
end


