class InspectionPlansController < ApplicationController
  def index
    @inspectionPlans = InspectionPlan.all
  end
 
  def show
    @inspectionPlan = InspectionPlan.find(params[:id])
  end
 
  def new
    @inspectionPlan = InspectionPlan.new
  end
 
  def edit
    @inspectionPlan = InspectionPlan.find(params[:id])
  end
 
  def create
    @inspectionPlan = InspectionPlan.new(inspectionPlan_params)
 
    if @inspectionPlan.save
      redirect_to inspectionPlans_path
    else
      render 'new'
    end
  end
 
  def update
    @inspectionPlan = InspectionPlan.find(params[:id])
 
    if @inspectionPlan.update(inspectionPlan_params)
      redirect_to inspectionPlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inspectionPlan = InspectionPlan.find(params[:id])
    @inspectionPlan.destroy
    redirect_to inspectionPlans_path
  end

 
  private
    def inspectionPlan_params
      params.require(:inspectionPlan).permit(:planNumber, :revision, :SamplingPlan, :Status)
    end
end