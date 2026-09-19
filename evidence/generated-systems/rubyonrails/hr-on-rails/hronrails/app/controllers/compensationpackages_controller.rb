class CompensationPackagesController < ApplicationController
  def index
    @compensationPackages = CompensationPackage.all
  end
 
  def show
    @compensationPackage = CompensationPackage.find(params[:id])
  end
 
  def new
    @compensationPackage = CompensationPackage.new
  end
 
  def edit
    @compensationPackage = CompensationPackage.find(params[:id])
  end
 
  def create
    @compensationPackage = CompensationPackage.new(compensationPackage_params)
 
    if @compensationPackage.save
      redirect_to compensationPackages_path
    else
      render 'new'
    end
  end
 
  def update
    @compensationPackage = CompensationPackage.find(params[:id])
 
    if @compensationPackage.update(compensationPackage_params)
      redirect_to compensationPackages_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @compensationPackage = CompensationPackage.find(params[:id])
    @compensationPackage.destroy
    redirect_to compensationPackages_path
  end

 
  private
    def compensationPackage_params
      params.require(:compensationPackage).permit(:effectiveFrom, :effectiveTo, :currency)
    end
end