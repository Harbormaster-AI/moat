class BusinessUnitsController < ApplicationController
  def index
    @businessUnits = BusinessUnit.all
  end
 
  def show
    @businessUnit = BusinessUnit.find(params[:id])
  end
 
  def new
    @businessUnit = BusinessUnit.new
  end
 
  def edit
    @businessUnit = BusinessUnit.find(params[:id])
  end
 
  def create
    @businessUnit = BusinessUnit.new(businessUnit_params)
 
    if @businessUnit.save
      redirect_to businessUnits_path
    else
      render 'new'
    end
  end
 
  def update
    @businessUnit = BusinessUnit.find(params[:id])
 
    if @businessUnit.update(businessUnit_params)
      redirect_to businessUnits_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @businessUnit = BusinessUnit.find(params[:id])
    @businessUnit.destroy
    redirect_to businessUnits_path
  end

 
  private
    def businessUnit_params
      params.require(:businessUnit).permit(:name, :leader)
    end
end