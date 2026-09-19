class PharmacysController < ApplicationController
  def index
    @pharmacys = Pharmacy.all
  end
 
  def show
    @pharmacy = Pharmacy.find(params[:id])
  end
 
  def new
    @pharmacy = Pharmacy.new
  end
 
  def edit
    @pharmacy = Pharmacy.find(params[:id])
  end
 
  def create
    @pharmacy = Pharmacy.new(pharmacy_params)
 
    if @pharmacy.save
      redirect_to pharmacys_path
    else
      render 'new'
    end
  end
 
  def update
    @pharmacy = Pharmacy.find(params[:id])
 
    if @pharmacy.update(pharmacy_params)
      redirect_to pharmacys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @pharmacy = Pharmacy.find(params[:id])
    @pharmacy.destroy
    redirect_to pharmacys_path
  end

 
  private
    def pharmacy_params
      params.require(:pharmacy).permit(:name)
    end
end