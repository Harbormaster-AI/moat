class ObservationsController < ApplicationController
  def index
    @observations = Observation.all
  end
 
  def show
    @observation = Observation.find(params[:id])
  end
 
  def new
    @observation = Observation.new
  end
 
  def edit
    @observation = Observation.find(params[:id])
  end
 
  def create
    @observation = Observation.new(observation_params)
 
    if @observation.save
      redirect_to observations_path
    else
      render 'new'
    end
  end
 
  def update
    @observation = Observation.find(params[:id])
 
    if @observation.update(observation_params)
      redirect_to observations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @observation = Observation.find(params[:id])
    @observation.destroy
    redirect_to observations_path
  end

 
  private
    def observation_params
      params.require(:observation).permit(:code, :value, :unit, :effectiveDateTime, :Interpretation)
    end
end