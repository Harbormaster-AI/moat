class MaintenancePlansController < ApplicationController
  def index
    @maintenancePlans = MaintenancePlan.all
  end
 
  def show
    @maintenancePlan = MaintenancePlan.find(params[:id])
  end
 
  def new
    @maintenancePlan = MaintenancePlan.new
  end
 
  def edit
    @maintenancePlan = MaintenancePlan.find(params[:id])
  end
 
  def create
    @maintenancePlan = MaintenancePlan.new(maintenancePlan_params)
 
    if @maintenancePlan.save
      redirect_to maintenancePlans_path
    else
      render 'new'
    end
  end
 
  def update
    @maintenancePlan = MaintenancePlan.find(params[:id])
 
    if @maintenancePlan.update(maintenancePlan_params)
      redirect_to maintenancePlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @maintenancePlan = MaintenancePlan.find(params[:id])
    @maintenancePlan.destroy
    redirect_to maintenancePlans_path
  end

 
  private
    def maintenancePlan_params
      params.require(:maintenancePlan).permit(:planNumber, :interval, :lastServiceDate, :Strategy)
    end
end