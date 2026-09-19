class AircraftProgramsController < ApplicationController
  def index
    @aircraftPrograms = AircraftProgram.all
  end
 
  def show
    @aircraftProgram = AircraftProgram.find(params[:id])
  end
 
  def new
    @aircraftProgram = AircraftProgram.new
  end
 
  def edit
    @aircraftProgram = AircraftProgram.find(params[:id])
  end
 
  def create
    @aircraftProgram = AircraftProgram.new(aircraftProgram_params)
 
    if @aircraftProgram.save
      redirect_to aircraftPrograms_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftProgram = AircraftProgram.find(params[:id])
 
    if @aircraftProgram.update(aircraftProgram_params)
      redirect_to aircraftPrograms_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftProgram = AircraftProgram.find(params[:id])
    @aircraftProgram.destroy
    redirect_to aircraftPrograms_path
  end

 
  private
    def aircraftProgram_params
      params.require(:aircraftProgram).permit(:name, :programCode, :entryIntoServiceYear, :Status)
    end
end