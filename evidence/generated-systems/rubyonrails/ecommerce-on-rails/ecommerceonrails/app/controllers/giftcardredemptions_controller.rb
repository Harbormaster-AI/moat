class GiftCardRedemptionsController < ApplicationController
  def index
    @giftCardRedemptions = GiftCardRedemption.all
  end
 
  def show
    @giftCardRedemption = GiftCardRedemption.find(params[:id])
  end
 
  def new
    @giftCardRedemption = GiftCardRedemption.new
  end
 
  def edit
    @giftCardRedemption = GiftCardRedemption.find(params[:id])
  end
 
  def create
    @giftCardRedemption = GiftCardRedemption.new(giftCardRedemption_params)
 
    if @giftCardRedemption.save
      redirect_to giftCardRedemptions_path
    else
      render 'new'
    end
  end
 
  def update
    @giftCardRedemption = GiftCardRedemption.find(params[:id])
 
    if @giftCardRedemption.update(giftCardRedemption_params)
      redirect_to giftCardRedemptions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @giftCardRedemption = GiftCardRedemption.find(params[:id])
    @giftCardRedemption.destroy
    redirect_to giftCardRedemptions_path
  end

 
  private
    def giftCardRedemption_params
      params.require(:giftCardRedemption).permit(:redeemedAt, :amount)
    end
end