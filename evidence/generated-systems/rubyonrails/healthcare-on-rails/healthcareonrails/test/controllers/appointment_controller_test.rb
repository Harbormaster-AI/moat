require "test_helper"

class AppointmentControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @appointment = appointments(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create appointment" do
    assert_difference("Appointment.count") do
      post appointments_url, params: { appointment: { appointmentDate:1.week.ago, reason:"test string for reason", Status:Appointment.Statuss[0], Priority:Appointment.Prioritys[0] } }
    end

    assert_redirected_to appointments_url
  end

 
  
  test "should destroy appointment" do
    assert_difference("Appointment.count", -1) do
      delete appointment_url(@appointment)
    end

    assert_redirected_to appointments_url
  end
  
end


