require "test_helper"

class DepartmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @department = departments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create department" do
    assert_difference("Department.count") do
      post departments_url, params: { department: { name:"test string for name", DepartmentType:Department.DepartmentTypes[0] } }
    end

    assert_redirected_to departments_url
  end

 
  
  test "should destroy department" do
    assert_difference("Department.count", -1) do
      delete department_url(@department)
    end

    assert_redirected_to departments_url
  end
  
end


