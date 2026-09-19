class MerchantsController < ApplicationController
  def index
    @merchants = Merchant.all
  end
 
  def show
    @merchant = Merchant.find(params[:id])
  end
 
  def new
    @merchant = Merchant.new
  end
 
  def edit
    @merchant = Merchant.find(params[:id])
  end
 
  def create
    @merchant = Merchant.new(merchant_params)
 
    if @merchant.save
      redirect_to merchants_path
    else
      render 'new'
    end
  end
 
  def update
    @merchant = Merchant.find(params[:id])
 
    if @merchant.update(merchant_params)
      redirect_to merchants_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @merchant = Merchant.find(params[:id])
    @merchant.destroy
    redirect_to merchants_path
  end

 
  private
    def merchant_params
      params.require(:merchant).permit(:name, :mcc, :url, :country, :settlementCurrency)
    end
end