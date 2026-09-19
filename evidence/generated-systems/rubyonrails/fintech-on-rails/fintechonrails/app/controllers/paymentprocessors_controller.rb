class PaymentProcessorsController < ApplicationController
  def index
    @paymentProcessors = PaymentProcessor.all
  end
 
  def show
    @paymentProcessor = PaymentProcessor.find(params[:id])
  end
 
  def new
    @paymentProcessor = PaymentProcessor.new
  end
 
  def edit
    @paymentProcessor = PaymentProcessor.find(params[:id])
  end
 
  def create
    @paymentProcessor = PaymentProcessor.new(paymentProcessor_params)
 
    if @paymentProcessor.save
      redirect_to paymentProcessors_path
    else
      render 'new'
    end
  end
 
  def update
    @paymentProcessor = PaymentProcessor.find(params[:id])
 
    if @paymentProcessor.update(paymentProcessor_params)
      redirect_to paymentProcessors_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @paymentProcessor = PaymentProcessor.find(params[:id])
    @paymentProcessor.destroy
    redirect_to paymentProcessors_path
  end

 
  private
    def paymentProcessor_params
      params.require(:paymentProcessor).permit(:name, :processorCode, :networkSupport)
    end
end