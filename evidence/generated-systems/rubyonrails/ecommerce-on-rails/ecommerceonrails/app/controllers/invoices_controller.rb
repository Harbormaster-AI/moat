class InvoicesController < ApplicationController
  def index
    @invoices = Invoice.all
  end
 
  def show
    @invoice = Invoice.find(params[:id])
  end
 
  def new
    @invoice = Invoice.new
  end
 
  def edit
    @invoice = Invoice.find(params[:id])
  end
 
  def create
    @invoice = Invoice.new(invoice_params)
 
    if @invoice.save
      redirect_to invoices_path
    else
      render 'new'
    end
  end
 
  def update
    @invoice = Invoice.find(params[:id])
 
    if @invoice.update(invoice_params)
      redirect_to invoices_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @invoice = Invoice.find(params[:id])
    @invoice.destroy
    redirect_to invoices_path
  end

 
  private
    def invoice_params
      params.require(:invoice).permit(:invoiceNumber, :issuedDate, :dueDate, :total, :Status)
    end
end