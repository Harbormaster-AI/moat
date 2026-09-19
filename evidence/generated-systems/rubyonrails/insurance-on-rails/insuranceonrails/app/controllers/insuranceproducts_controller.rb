class InsuranceProductsController < ApplicationController
  def index
    @insuranceProducts = InsuranceProduct.all
  end
 
  def show
    @insuranceProduct = InsuranceProduct.find(params[:id])
  end
 
  def new
    @insuranceProduct = InsuranceProduct.new
  end
 
  def edit
    @insuranceProduct = InsuranceProduct.find(params[:id])
  end
 
  def create
    @insuranceProduct = InsuranceProduct.new(insuranceProduct_params)
 
    if @insuranceProduct.save
      redirect_to insuranceProducts_path
    else
      render 'new'
    end
  end
 
  def update
    @insuranceProduct = InsuranceProduct.find(params[:id])
 
    if @insuranceProduct.update(insuranceProduct_params)
      redirect_to insuranceProducts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @insuranceProduct = InsuranceProduct.find(params[:id])
    @insuranceProduct.destroy
    redirect_to insuranceProducts_path
  end

 
  private
    def insuranceProduct_params
      params.require(:insuranceProduct).permit(:name, :productCode, :LineOfBusiness)
    end
end