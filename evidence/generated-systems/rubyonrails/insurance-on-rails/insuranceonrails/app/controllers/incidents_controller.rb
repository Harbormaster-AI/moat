class IncidentsController < ApplicationController
  def index
    @incidents = Incident.all
  end
 
  def show
    @incident = Incident.find(params[:id])
  end
 
  def new
    @incident = Incident.new
  end
 
  def edit
    @incident = Incident.find(params[:id])
  end
 
  def create
    @incident = Incident.new(incident_params)
 
    if @incident.save
      redirect_to incidents_path
    else
      render 'new'
    end
  end
 
  def update
    @incident = Incident.find(params[:id])
 
    if @incident.update(incident_params)
      redirect_to incidents_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @incident = Incident.find(params[:id])
    @incident.destroy
    redirect_to incidents_path
  end

 
  private
    def incident_params
      params.require(:incident).permit(:location, :description, :IncidentType)
    end
end