class CouponRedemptionsController < ApplicationController
  def index
    @couponRedemptions = CouponRedemption.all
  end
 
  def show
    @couponRedemption = CouponRedemption.find(params[:id])
  end
 
  def new
    @couponRedemption = CouponRedemption.new
  end
 
  def edit
    @couponRedemption = CouponRedemption.find(params[:id])
  end
 
  def create
    @couponRedemption = CouponRedemption.new(couponRedemption_params)
 
    if @couponRedemption.save
      redirect_to couponRedemptions_path
    else
      render 'new'
    end
  end
 
  def update
    @couponRedemption = CouponRedemption.find(params[:id])
 
    if @couponRedemption.update(couponRedemption_params)
      redirect_to couponRedemptions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @couponRedemption = CouponRedemption.find(params[:id])
    @couponRedemption.destroy
    redirect_to couponRedemptions_path
  end

 
  private
    def couponRedemption_params
      params.require(:couponRedemption).permit(:redeemedAt)
    end
end