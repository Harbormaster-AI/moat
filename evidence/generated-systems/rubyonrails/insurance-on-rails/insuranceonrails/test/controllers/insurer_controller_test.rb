require "test_helper"

class InsurerControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @insurer = insurers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create insurer" do
    assert_difference("Insurer.count") do
      post insurers_url, params: { insurer: { name:"test string for name", legalName:"test string for legalName", domicileCountry:"test string for domicileCountry", naicNumber:"test string for naicNumber", website:"test string for website" } }
    end

    assert_redirected_to insurers_url
  end

 
  
  test "should destroy insurer" do
    assert_difference("Insurer.count", -1) do
      delete insurer_url(@insurer)
    end

    assert_redirected_to insurers_url
  end
  
end


