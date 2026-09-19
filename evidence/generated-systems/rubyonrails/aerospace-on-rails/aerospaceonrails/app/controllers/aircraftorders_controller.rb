class AircraftOrdersController < ApplicationController
  def index
    @aircraftOrders = AircraftOrder.all
  end
 
  def show
    @aircraftOrder = AircraftOrder.find(params[:id])
  end
 
  def new
    @aircraftOrder = AircraftOrder.new
  end
 
  def edit
    @aircraftOrder = AircraftOrder.find(params[:id])
  end
 
  def create
    @aircraftOrder = AircraftOrder.new(aircraftOrder_params)
 
    if @aircraftOrder.save
      redirect_to aircraftOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @aircraftOrder = AircraftOrder.find(params[:id])
 
    if @aircraftOrder.update(aircraftOrder_params)
      redirect_to aircraftOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aircraftOrder = AircraftOrder.find(params[:id])
    @aircraftOrder.destroy
    redirect_to aircraftOrders_path
  end

 
  private
    def aircraftOrder_params
      params.require(:aircraftOrder).permit(:orderNumber, :totalAmount, :Status)
    end
end