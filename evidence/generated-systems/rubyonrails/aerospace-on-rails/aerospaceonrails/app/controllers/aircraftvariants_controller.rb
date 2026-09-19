class AircraftVariantsController < ApplicationController
  def index
    @aircraftVariants = AircraftVariant.all
  end
 
  def show
    @aircraftVariant = AircraftVariant.find(params[:id])
  end
 
  def new
    @aircraftVariant = AircraftVariant.new
  end
 
  def edit
    @aircraftVariant = AircraftVariant.find(params[:id])
  end
 
  def create
    @aircraftVariant = AircraftVariant.new(aircraftVariant_params)
 
    if @aircraftVariant.save
      redirect_to aircraftVariants_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftVariant = AircraftVariant.find(params[:id])
 
    if @aircraftVariant.update(aircraftVariant_params)
      redirect_to aircraftVariants_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftVariant = AircraftVariant.find(params[:id])
    @aircraftVariant.destroy
    redirect_to aircraftVariants_path
  end

 
  private
    def aircraftVariant_params
      params.require(:aircraftVariant).permit(:variantCode, :rangeNm, :maxTakeoffWeightKg)
    end
end