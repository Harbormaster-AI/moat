class ProductOfferingsController < ApplicationController
  def index
    @productOfferings = ProductOffering.all
  end
 
  def show
    @productOffering = ProductOffering.find(params[:id])
  end
 
  def new
    @productOffering = ProductOffering.new
  end
 
  def edit
    @productOffering = ProductOffering.find(params[:id])
  end
 
  def create
    @productOffering = ProductOffering.new(productOffering_params)
 
    if @productOffering.save
      redirect_to productOfferings_path
    else
      render 'new'
    end
  end
 
  def update
    @productOffering = ProductOffering.find(params[:id])
 
    if @productOffering.update(productOffering_params)
      redirect_to productOfferings_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productOffering = ProductOffering.find(params[:id])
    @productOffering.destroy
    redirect_to productOfferings_path
  end

 
  private
    def productOffering_params
      params.require(:productOffering).permit(:name, :productCode, :Category)
    end
end