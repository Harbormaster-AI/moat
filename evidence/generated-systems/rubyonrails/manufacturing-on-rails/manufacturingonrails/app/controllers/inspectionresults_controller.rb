class InspectionResultsController < ApplicationController
  def index
    @inspectionResults = InspectionResult.all
  end
 
  def show
    @inspectionResult = InspectionResult.find(params[:id])
  end
 
  def new
    @inspectionResult = InspectionResult.new
  end
 
  def edit
    @inspectionResult = InspectionResult.find(params[:id])
  end
 
  def create
    @inspectionResult = InspectionResult.new(inspectionResult_params)
 
    if @inspectionResult.save
      redirect_to inspectionResults_path
    else
      render 'new'
    end
  end
 
  def update
    @inspectionResult = InspectionResult.find(params[:id])
 
    if @inspectionResult.update(inspectionResult_params)
      redirect_to inspectionResults_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inspectionResult = InspectionResult.find(params[:id])
    @inspectionResult.destroy
    redirect_to inspectionResults_path
  end

 
  private
    def inspectionResult_params
      params.require(:inspectionResult).permit(:resultValue, :recordedOn, :notes, :ResultStatus)
    end
end