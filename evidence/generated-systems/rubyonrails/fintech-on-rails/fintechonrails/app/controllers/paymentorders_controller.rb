class PaymentOrdersController < ApplicationController
  def index
    @paymentOrders = PaymentOrder.all
  end
 
  def show
    @paymentOrder = PaymentOrder.find(params[:id])
  end
 
  def new
    @paymentOrder = PaymentOrder.new
  end
 
  def edit
    @paymentOrder = PaymentOrder.find(params[:id])
  end
 
  def create
    @paymentOrder = PaymentOrder.new(paymentOrder_params)
 
    if @paymentOrder.save
      redirect_to paymentOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @paymentOrder = PaymentOrder.find(params[:id])
 
    if @paymentOrder.update(paymentOrder_params)
      redirect_to paymentOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @paymentOrder = PaymentOrder.find(params[:id])
    @paymentOrder.destroy
    redirect_to paymentOrders_path
  end

 
  private
    def paymentOrder_params
      params.require(:paymentOrder).permit(:orderReference, :requestedExecutionDate, :purpose, :PaymentMethod, :Status, :Priority)
    end
end