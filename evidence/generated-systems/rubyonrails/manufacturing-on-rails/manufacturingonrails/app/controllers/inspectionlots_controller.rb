class InspectionLotsController < ApplicationController
  def index
    @inspectionLots = InspectionLot.all
  end
 
  def show
    @inspectionLot = InspectionLot.find(params[:id])
  end
 
  def new
    @inspectionLot = InspectionLot.new
  end
 
  def edit
    @inspectionLot = InspectionLot.find(params[:id])
  end
 
  def create
    @inspectionLot = InspectionLot.new(inspectionLot_params)
 
    if @inspectionLot.save
      redirect_to inspectionLots_path
    else
      render 'new'
    end
  end
 
  def update
    @inspectionLot = InspectionLot.find(params[:id])
 
    if @inspectionLot.update(inspectionLot_params)
      redirect_to inspectionLots_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inspectionLot = InspectionLot.find(params[:id])
    @inspectionLot.destroy
    redirect_to inspectionLots_path
  end

 
  private
    def inspectionLot_params
      params.require(:inspectionLot).permit(:lotNumber, :quantity, :sampleSize, :createdOn, :InspectionType, :Status)
    end
end