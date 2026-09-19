class CouponsController < ApplicationController
  def index
    @coupons = Coupon.all
  end
 
  def show
    @coupon = Coupon.find(params[:id])
  end
 
  def new
    @coupon = Coupon.new
  end
 
  def edit
    @coupon = Coupon.find(params[:id])
  end
 
  def create
    @coupon = Coupon.new(coupon_params)
 
    if @coupon.save
      redirect_to coupons_path
    else
      render 'new'
    end
  end
 
  def update
    @coupon = Coupon.find(params[:id])
 
    if @coupon.update(coupon_params)
      redirect_to coupons_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @coupon = Coupon.find(params[:id])
    @coupon.destroy
    redirect_to coupons_path
  end

 
  private
    def coupon_params
      params.require(:coupon).permit(:code, :usageLimit, :perCustomerLimit, :expirationDate, :Status)
    end
end