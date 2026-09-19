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
      post employees_url, params: { employee: { employeeNumber:"test string for employeeNumber", name:"test value", workEmail:"test value", workPhone:"test value", dateOfHire:1.week.ago, nationalId:"test value", Status:Employee.Statuss[0] } }
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


