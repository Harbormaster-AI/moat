class FXDealsController < ApplicationController
  def index
    @fXDeals = FXDeal.all
  end
 
  def show
    @fXDeal = FXDeal.find(params[:id])
  end
 
  def new
    @fXDeal = FXDeal.new
  end
 
  def edit
    @fXDeal = FXDeal.find(params[:id])
  end
 
  def create
    @fXDeal = FXDeal.new(fXDeal_params)
 
    if @fXDeal.save
      redirect_to fXDeals_path
    else
      render 'new'
    end
  end
 
  def update
    @fXDeal = FXDeal.find(params[:id])
 
    if @fXDeal.update(fXDeal_params)
      redirect_to fXDeals_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @fXDeal = FXDeal.find(params[:id])
    @fXDeal.destroy
    redirect_to fXDeals_path
  end

 
  private
    def fXDeal_params
      params.require(:fXDeal).permit(:dealReference, :baseCurrency, :quoteCurrency, :rate, :amount, :settlementDate, :Status)
    end
end