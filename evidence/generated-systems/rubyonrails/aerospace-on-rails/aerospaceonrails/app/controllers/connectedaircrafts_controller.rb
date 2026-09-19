class ConnectedAircraftsController < ApplicationController
  def index
    @connectedAircrafts = ConnectedAircraft.all
  end
 
  def show
    @connectedAircraft = ConnectedAircraft.find(params[:id])
  end
 
  def new
    @connectedAircraft = ConnectedAircraft.new
  end
 
  def edit
    @connectedAircraft = ConnectedAircraft.find(params[:id])
  end
 
  def create
    @connectedAircraft = ConnectedAircraft.new(connectedAircraft_params)
 
    if @connectedAircraft.save
      redirect_to connectedAircrafts_path
    else
      render 'new'
    end
  end
 
  def update
    @connectedAircraft = ConnectedAircraft.find(params[:id])
 
    if @connectedAircraft.update(connectedAircraft_params)
      redirect_to connectedAircrafts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @connectedAircraft = ConnectedAircraft.find(params[:id])
    @connectedAircraft.destroy
    redirect_to connectedAircrafts_path
  end

 
  private
    def connectedAircraft_params
      params.require(:connectedAircraft).permit(:communicationsProvider, :ConnectivityStatus)
    end
end