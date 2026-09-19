class PaymentContractsController < ApplicationController
  def index
    @paymentContracts = PaymentContract.all
  end
 
  def show
    @paymentContract = PaymentContract.find(params[:id])
  end
 
  def new
    @paymentContract = PaymentContract.new
  end
 
  def edit
    @paymentContract = PaymentContract.find(params[:id])
  end
 
  def create
    @paymentContract = PaymentContract.new(paymentContract_params)
 
    if @paymentContract.save
      redirect_to paymentContracts_path
    else
      render 'new'
    end
  end
 
  def update
    @paymentContract = PaymentContract.find(params[:id])
 
    if @paymentContract.update(paymentContract_params)
      redirect_to paymentContracts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @paymentContract = PaymentContract.find(params[:id])
    @paymentContract.destroy
    redirect_to paymentContracts_path
  end

 
  private
    def paymentContract_params
      params.require(:paymentContract).permit(:contractNumber, :pricingPlanCode, :Status)
    end
end