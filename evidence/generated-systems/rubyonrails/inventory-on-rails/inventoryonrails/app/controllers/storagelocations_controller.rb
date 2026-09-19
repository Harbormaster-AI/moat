class StorageLocationsController < ApplicationController
  def index
    @storageLocations = StorageLocation.all
  end
 
  def show
    @storageLocation = StorageLocation.find(params[:id])
  end
 
  def new
    @storageLocation = StorageLocation.new
  end
 
  def edit
    @storageLocation = StorageLocation.find(params[:id])
  end
 
  def create
    @storageLocation = StorageLocation.new(storageLocation_params)
 
    if @storageLocation.save
      redirect_to storageLocations_path
    else
      render 'new'
    end
  end
 
  def update
    @storageLocation = StorageLocation.find(params[:id])
 
    if @storageLocation.update(storageLocation_params)
      redirect_to storageLocations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @storageLocation = StorageLocation.find(params[:id])
    @storageLocation.destroy
    redirect_to storageLocations_path
  end

 
  private
    def storageLocation_params
      params.require(:storageLocation).permit(:code, :temperatureControlled, :capacity, :capacityUnit, :LocationType)
    end
end