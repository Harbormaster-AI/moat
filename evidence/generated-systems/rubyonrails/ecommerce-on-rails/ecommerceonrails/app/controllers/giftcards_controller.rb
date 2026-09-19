class GiftCardsController < ApplicationController
  def index
    @giftCards = GiftCard.all
  end
 
  def show
    @giftCard = GiftCard.find(params[:id])
  end
 
  def new
    @giftCard = GiftCard.new
  end
 
  def edit
    @giftCard = GiftCard.find(params[:id])
  end
 
  def create
    @giftCard = GiftCard.new(giftCard_params)
 
    if @giftCard.save
      redirect_to giftCards_path
    else
      render 'new'
    end
  end
 
  def update
    @giftCard = GiftCard.find(params[:id])
 
    if @giftCard.update(giftCard_params)
      redirect_to giftCards_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @giftCard = GiftCard.find(params[:id])
    @giftCard.destroy
    redirect_to giftCards_path
  end

 
  private
    def giftCard_params
      params.require(:giftCard).permit(:code, :balance, :expirationDate, :Status)
    end
end