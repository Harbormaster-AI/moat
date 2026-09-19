class ProductPricingsController < ApplicationController
  def index
    @productPricings = ProductPricing.all
  end
 
  def show
    @productPricing = ProductPricing.find(params[:id])
  end
 
  def new
    @productPricing = ProductPricing.new
  end
 
  def edit
    @productPricing = ProductPricing.find(params[:id])
  end
 
  def create
    @productPricing = ProductPricing.new(productPricing_params)
 
    if @productPricing.save
      redirect_to productPricings_path
    else
      render 'new'
    end
  end
 
  def update
    @productPricing = ProductPricing.find(params[:id])
 
    if @productPricing.update(productPricing_params)
      redirect_to productPricings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productPricing = ProductPricing.find(params[:id])
    @productPricing.destroy
    redirect_to productPricings_path
  end

 
  private
    def productPricing_params
      params.require(:productPricing).permit(:listPrice, :salePrice, :validFrom, :validTo)
    end
end