class FlightHealthEventsController < ApplicationController
  def index
    @flightHealthEvents = FlightHealthEvent.all
  end
 
  def show
    @flightHealthEvent = FlightHealthEvent.find(params[:id])
  end
 
  def new
    @flightHealthEvent = FlightHealthEvent.new
  end
 
  def edit
    @flightHealthEvent = FlightHealthEvent.find(params[:id])
  end
 
  def create
    @flightHealthEvent = FlightHealthEvent.new(flightHealthEvent_params)
 
    if @flightHealthEvent.save
      redirect_to flightHealthEvents_path
    else
      render 'new'
    end
  end
 
  def update
    @flightHealthEvent = FlightHealthEvent.find(params[:id])
 
    if @flightHealthEvent.update(flightHealthEvent_params)
      redirect_to flightHealthEvents_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @flightHealthEvent = FlightHealthEvent.find(params[:id])
    @flightHealthEvent.destroy
    redirect_to flightHealthEvents_path
  end

 
  private
    def flightHealthEvent_params
      params.require(:flightHealthEvent).permit(:eventCode, :Severity)
    end
end