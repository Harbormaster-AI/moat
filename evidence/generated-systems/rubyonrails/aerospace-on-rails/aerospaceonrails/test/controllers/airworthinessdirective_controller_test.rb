require "test_helper"

class AirworthinessDirectiveControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @airworthinessDirective = airworthinessDirectives(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create airworthinessDirective" do
    assert_difference("AirworthinessDirective.count") do
      post airworthinessDirectives_url, params: { airworthinessDirective: { directiveNumber:"test string for directiveNumber", title:"test string for title" } }
    end

    assert_redirected_to airworthinessDirectives_url
  end

 
  
  test "should destroy airworthinessDirective" do
    assert_difference("AirworthinessDirective.count", -1) do
      delete airworthinessDirective_url(@airworthinessDirective)
    end

    assert_redirected_to airworthinessDirectives_url
  end
  
end


