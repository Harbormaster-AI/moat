class InspectionCharacteristicsController < ApplicationController
  def index
    @inspectionCharacteristics = InspectionCharacteristic.all
  end
 
  def show
    @inspectionCharacteristic = InspectionCharacteristic.find(params[:id])
  end
 
  def new
    @inspectionCharacteristic = InspectionCharacteristic.new
  end
 
  def edit
    @inspectionCharacteristic = InspectionCharacteristic.find(params[:id])
  end
 
  def create
    @inspectionCharacteristic = InspectionCharacteristic.new(inspectionCharacteristic_params)
 
    if @inspectionCharacteristic.save
      redirect_to inspectionCharacteristics_path
    else
      render 'new'
    end
  end
 
  def update
    @inspectionCharacteristic = InspectionCharacteristic.find(params[:id])
 
    if @inspectionCharacteristic.update(inspectionCharacteristic_params)
      redirect_to inspectionCharacteristics_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inspectionCharacteristic = InspectionCharacteristic.find(params[:id])
    @inspectionCharacteristic.destroy
    redirect_to inspectionCharacteristics_path
  end

 
  private
    def inspectionCharacteristic_params
      params.require(:inspectionCharacteristic).permit(:characteristicCode, :name, :lowerSpecLimit, :upperSpecLimit, :target, :MeasurementType)
    end
end