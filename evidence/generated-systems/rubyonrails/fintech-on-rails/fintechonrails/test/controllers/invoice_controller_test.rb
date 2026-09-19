require "test_helper"

class InvoiceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @invoice = invoices(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create invoice" do
    assert_difference("Invoice.count") do
      post invoices_url, params: { invoice: { invoiceNumber:"test string for invoiceNumber", issueDate:1.week.ago, dueDate:1.week.ago, total:"test value", currency:"test string for currency", Status:Invoice.Statuss[0] } }
    end

    assert_redirected_to invoices_url
  end

 
  
  test "should destroy invoice" do
    assert_difference("Invoice.count", -1) do
      delete invoice_url(@invoice)
    end

    assert_redirected_to invoices_url
  end
  
end


