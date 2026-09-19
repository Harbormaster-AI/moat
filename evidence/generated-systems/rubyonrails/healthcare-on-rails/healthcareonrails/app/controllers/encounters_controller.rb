class EncountersController < ApplicationController
  def index
    @encounters = Encounter.all
  end
 
  def show
    @encounter = Encounter.find(params[:id])
  end
 
  def new
    @encounter = Encounter.new
  end
 
  def edit
    @encounter = Encounter.find(params[:id])
  end
 
  def create
    @encounter = Encounter.new(encounter_params)
 
    if @encounter.save
      redirect_to encounters_path
    else
      render 'new'
    end
  end
 
  def update
    @encounter = Encounter.find(params[:id])
 
    if @encounter.update(encounter_params)
      redirect_to encounters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @encounter = Encounter.find(params[:id])
    @encounter.destroy
    redirect_to encounters_path
  end

 
  private
    def encounter_params
      params.require(:encounter).permit(:encounterNumber, :startDateTime, :endDateTime, :Status, :EncounterType)
    end
end