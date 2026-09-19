class PromotionsController < ApplicationController
  def index
    @promotions = Promotion.all
  end
 
  def show
    @promotion = Promotion.find(params[:id])
  end
 
  def new
    @promotion = Promotion.new
  end
 
  def edit
    @promotion = Promotion.find(params[:id])
  end
 
  def create
    @promotion = Promotion.new(promotion_params)
 
    if @promotion.save
      redirect_to promotions_path
    else
      render 'new'
    end
  end
 
  def update
    @promotion = Promotion.find(params[:id])
 
    if @promotion.update(promotion_params)
      redirect_to promotions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @promotion = Promotion.find(params[:id])
    @promotion.destroy
    redirect_to promotions_path
  end

 
  private
    def promotion_params
      params.require(:promotion).permit(:name, :code, :value, :startDate, :endDate, :asStackable, :maxRedemptions, :PromotionType, :DiscountType)
    end
end