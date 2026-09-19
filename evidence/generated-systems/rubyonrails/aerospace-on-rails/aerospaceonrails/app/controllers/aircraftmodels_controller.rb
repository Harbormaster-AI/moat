class AircraftModelsController < ApplicationController
  def index
    @aircraftModels = AircraftModel.all
  end
 
  def show
    @aircraftModel = AircraftModel.find(params[:id])
  end
 
  def new
    @aircraftModel = AircraftModel.new
  end
 
  def edit
    @aircraftModel = AircraftModel.find(params[:id])
  end
 
  def create
    @aircraftModel = AircraftModel.new(aircraftModel_params)
 
    if @aircraftModel.save
      redirect_to aircraftModels_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftModel = AircraftModel.find(params[:id])
 
    if @aircraftModel.update(aircraftModel_params)
      redirect_to aircraftModels_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftModel = AircraftModel.find(params[:id])
    @aircraftModel.destroy
    redirect_to aircraftModels_path
  end

 
  private
    def aircraftModel_params
      params.require(:aircraftModel).permit(:name, :modelDesignation, :AircraftType)
    end
end