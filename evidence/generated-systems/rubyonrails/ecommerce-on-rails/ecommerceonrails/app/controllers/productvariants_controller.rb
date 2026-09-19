class ProductVariantsController < ApplicationController
  def index
    @productVariants = ProductVariant.all
  end
 
  def show
    @productVariant = ProductVariant.find(params[:id])
  end
 
  def new
    @productVariant = ProductVariant.new
  end
 
  def edit
    @productVariant = ProductVariant.find(params[:id])
  end
 
  def create
    @productVariant = ProductVariant.new(productVariant_params)
 
    if @productVariant.save
      redirect_to productVariants_path
    else
      render 'new'
    end
  end
 
  def update
    @productVariant = ProductVariant.find(params[:id])
 
    if @productVariant.update(productVariant_params)
      redirect_to productVariants_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productVariant = ProductVariant.find(params[:id])
    @productVariant.destroy
    redirect_to productVariants_path
  end

 
  private
    def productVariant_params
      params.require(:productVariant).permit(:sku, :barcode, :title, :weight, :requiresShipping, :WeightUnit)
    end
end