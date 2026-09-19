class PaymentCardsController < ApplicationController
  def index
    @paymentCards = PaymentCard.all
  end
 
  def show
    @paymentCard = PaymentCard.find(params[:id])
  end
 
  def new
    @paymentCard = PaymentCard.new
  end
 
  def edit
    @paymentCard = PaymentCard.find(params[:id])
  end
 
  def create
    @paymentCard = PaymentCard.new(paymentCard_params)
 
    if @paymentCard.save
      redirect_to paymentCards_path
    else
      render 'new'
    end
  end
 
  def update
    @paymentCard = PaymentCard.find(params[:id])
 
    if @paymentCard.update(paymentCard_params)
      redirect_to paymentCards_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @paymentCard = PaymentCard.find(params[:id])
    @paymentCard.destroy
    redirect_to paymentCards_path
  end

 
  private
    def paymentCard_params
      params.require(:paymentCard).permit(:cardToken, :maskedPan, :expiryMonth, :expiryYear, :cardholderName, :Scheme, :Status)
    end
end