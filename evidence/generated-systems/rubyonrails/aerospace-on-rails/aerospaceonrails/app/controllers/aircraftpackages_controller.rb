class AircraftPackagesController < ApplicationController
  def index
    @aircraftPackages = AircraftPackage.all
  end
 
  def show
    @aircraftPackage = AircraftPackage.find(params[:id])
  end
 
  def new
    @aircraftPackage = AircraftPackage.new
  end
 
  def edit
    @aircraftPackage = AircraftPackage.find(params[:id])
  end
 
  def create
    @aircraftPackage = AircraftPackage.new(aircraftPackage_params)
 
    if @aircraftPackage.save
      redirect_to aircraftPackages_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftPackage = AircraftPackage.find(params[:id])
 
    if @aircraftPackage.update(aircraftPackage_params)
      redirect_to aircraftPackages_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftPackage = AircraftPackage.find(params[:id])
    @aircraftPackage.destroy
    redirect_to aircraftPackages_path
  end

 
  private
    def aircraftPackage_params
      params.require(:aircraftPackage).permit(:name, :PackageType)
    end
end