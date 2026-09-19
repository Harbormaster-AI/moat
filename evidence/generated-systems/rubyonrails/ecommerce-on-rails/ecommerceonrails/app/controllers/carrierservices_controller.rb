class CarrierServicesController < ApplicationController
  def index
    @carrierServices = CarrierService.all
  end
 
  def show
    @carrierService = CarrierService.find(params[:id])
  end
 
  def new
    @carrierService = CarrierService.new
  end
 
  def edit
    @carrierService = CarrierService.find(params[:id])
  end
 
  def create
    @carrierService = CarrierService.new(carrierService_params)
 
    if @carrierService.save
      redirect_to carrierServices_path
    else
      render 'new'
    end
  end
 
  def update
    @carrierService = CarrierService.find(params[:id])
 
    if @carrierService.update(carrierService_params)
      redirect_to carrierServices_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @carrierService = CarrierService.find(params[:id])
    @carrierService.destroy
    redirect_to carrierServices_path
  end

 
  private
    def carrierService_params
      params.require(:carrierService).permit(:name, :code, :Carrier, :ServiceLevel)
    end
end