class AircraftFamilysController < ApplicationController
  def index
    @aircraftFamilys = AircraftFamily.all
  end
 
  def show
    @aircraftFamily = AircraftFamily.find(params[:id])
  end
 
  def new
    @aircraftFamily = AircraftFamily.new
  end
 
  def edit
    @aircraftFamily = AircraftFamily.find(params[:id])
  end
 
  def create
    @aircraftFamily = AircraftFamily.new(aircraftFamily_params)
 
    if @aircraftFamily.save
      redirect_to aircraftFamilys_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftFamily = AircraftFamily.find(params[:id])
 
    if @aircraftFamily.update(aircraftFamily_params)
      redirect_to aircraftFamilys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftFamily = AircraftFamily.find(params[:id])
    @aircraftFamily.destroy
    redirect_to aircraftFamilys_path
  end

 
  private
    def aircraftFamily_params
      params.require(:aircraftFamily).permit(:name, :familyCode)
    end
end