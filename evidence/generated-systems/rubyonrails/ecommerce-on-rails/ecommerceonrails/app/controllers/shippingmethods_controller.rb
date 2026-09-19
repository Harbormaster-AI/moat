class ShippingMethodsController < ApplicationController
  def index
    @shippingMethods = ShippingMethod.all
  end
 
  def show
    @shippingMethod = ShippingMethod.find(params[:id])
  end
 
  def new
    @shippingMethod = ShippingMethod.new
  end
 
  def edit
    @shippingMethod = ShippingMethod.find(params[:id])
  end
 
  def create
    @shippingMethod = ShippingMethod.new(shippingMethod_params)
 
    if @shippingMethod.save
      redirect_to shippingMethods_path
    else
      render 'new'
    end
  end
 
  def update
    @shippingMethod = ShippingMethod.find(params[:id])
 
    if @shippingMethod.update(shippingMethod_params)
      redirect_to shippingMethods_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @shippingMethod = ShippingMethod.find(params[:id])
    @shippingMethod.destroy
    redirect_to shippingMethods_path
  end

 
  private
    def shippingMethod_params
      params.require(:shippingMethod).permit(:name, :flatRate, :estimatedDays, :asActive, :MethodType)
    end
end