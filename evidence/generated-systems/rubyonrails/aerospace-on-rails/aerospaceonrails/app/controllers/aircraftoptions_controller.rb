class AircraftOptionsController < ApplicationController
  def index
    @aircraftOptions = AircraftOption.all
  end
 
  def show
    @aircraftOption = AircraftOption.find(params[:id])
  end
 
  def new
    @aircraftOption = AircraftOption.new
  end
 
  def edit
    @aircraftOption = AircraftOption.find(params[:id])
  end
 
  def create
    @aircraftOption = AircraftOption.new(aircraftOption_params)
 
    if @aircraftOption.save
      redirect_to aircraftOptions_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftOption = AircraftOption.find(params[:id])
 
    if @aircraftOption.update(aircraftOption_params)
      redirect_to aircraftOptions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftOption = AircraftOption.find(params[:id])
    @aircraftOption.destroy
    redirect_to aircraftOptions_path
  end

 
  private
    def aircraftOption_params
      params.require(:aircraftOption).permit(:code, :name, :OptionCategory)
    end
end