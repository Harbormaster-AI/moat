class CostCentersController < ApplicationController
  def index
    @costCenters = CostCenter.all
  end
 
  def show
    @costCenter = CostCenter.find(params[:id])
  end
 
  def new
    @costCenter = CostCenter.new
  end
 
  def edit
    @costCenter = CostCenter.find(params[:id])
  end
 
  def create
    @costCenter = CostCenter.new(costCenter_params)
 
    if @costCenter.save
      redirect_to costCenters_path
    else
      render 'new'
    end
  end
 
  def update
    @costCenter = CostCenter.find(params[:id])
 
    if @costCenter.update(costCenter_params)
      redirect_to costCenters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @costCenter = CostCenter.find(params[:id])
    @costCenter.destroy
    redirect_to costCenters_path
  end

 
  private
    def costCenter_params
      params.require(:costCenter).permit(:code, :name)
    end
end