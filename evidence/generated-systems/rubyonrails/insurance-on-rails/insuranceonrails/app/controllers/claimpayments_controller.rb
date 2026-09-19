class ClaimPaymentsController < ApplicationController
  def index
    @claimPayments = ClaimPayment.all
  end
 
  def show
    @claimPayment = ClaimPayment.find(params[:id])
  end
 
  def new
    @claimPayment = ClaimPayment.new
  end
 
  def edit
    @claimPayment = ClaimPayment.find(params[:id])
  end
 
  def create
    @claimPayment = ClaimPayment.new(claimPayment_params)
 
    if @claimPayment.save
      redirect_to claimPayments_path
    else
      render 'new'
    end
  end
 
  def update
    @claimPayment = ClaimPayment.find(params[:id])
 
    if @claimPayment.update(claimPayment_params)
      redirect_to claimPayments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @claimPayment = ClaimPayment.find(params[:id])
    @claimPayment.destroy
    redirect_to claimPayments_path
  end

 
  private
    def claimPayment_params
      params.require(:claimPayment).permit(:paymentNumber, :amount, :paymentDate, :PayeeType, :Method, :Status)
    end
end