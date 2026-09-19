require "test_helper"

class SalaryComponentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @salaryComponent = salaryComponents(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create salaryComponent" do
    assert_difference("SalaryComponent.count") do
      post salaryComponents_url, params: { salaryComponent: { amount:"test value", recurring:true, ComponentType:SalaryComponent.ComponentTypes[0] } }
    end

    assert_redirected_to salaryComponents_url
  end

 
  
  test "should destroy salaryComponent" do
    assert_difference("SalaryComponent.count", -1) do
      delete salaryComponent_url(@salaryComponent)
    end

    assert_redirected_to salaryComponents_url
  end
  
end


