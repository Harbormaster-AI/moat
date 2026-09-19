class PaymentProvidersController < ApplicationController
  def index
    @paymentProviders = PaymentProvider.all
  end
 
  def show
    @paymentProvider = PaymentProvider.find(params[:id])
  end
 
  def new
    @paymentProvider = PaymentProvider.new
  end
 
  def edit
    @paymentProvider = PaymentProvider.find(params[:id])
  end
 
  def create
    @paymentProvider = PaymentProvider.new(paymentProvider_params)
 
    if @paymentProvider.save
      redirect_to paymentProviders_path
    else
      render 'new'
    end
  end
 
  def update
    @paymentProvider = PaymentProvider.find(params[:id])
 
    if @paymentProvider.update(paymentProvider_params)
      redirect_to paymentProviders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @paymentProvider = PaymentProvider.find(params[:id])
    @paymentProvider.destroy
    redirect_to paymentProviders_path
  end

 
  private
    def paymentProvider_params
      params.require(:paymentProvider).permit(:name, :enabled, :merchantAccountId, :ProviderType)
    end
end