require "test_helper"

class EmployeeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @employee = employees(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create employee" do
    assert_difference("Employee.count") do
      post employees_url, params: { employee: { firstName:"test string for firstName", lastName:"test string for lastName", Role:Employee.Roles[0], SkillLevel:Employee.SkillLevels[0] } }
    end

    assert_redirected_to employees_url
  end

 
  
  test "should destroy employee" do
    assert_difference("Employee.count", -1) do
      delete employee_url(@employee)
    end

    assert_redirected_to employees_url
  end
  
end


