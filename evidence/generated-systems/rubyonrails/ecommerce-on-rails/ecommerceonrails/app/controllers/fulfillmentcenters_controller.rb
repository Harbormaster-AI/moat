class FulfillmentCentersController < ApplicationController
  def index
    @fulfillmentCenters = FulfillmentCenter.all
  end
 
  def show
    @fulfillmentCenter = FulfillmentCenter.find(params[:id])
  end
 
  def new
    @fulfillmentCenter = FulfillmentCenter.new
  end
 
  def edit
    @fulfillmentCenter = FulfillmentCenter.find(params[:id])
  end
 
  def create
    @fulfillmentCenter = FulfillmentCenter.new(fulfillmentCenter_params)
 
    if @fulfillmentCenter.save
      redirect_to fulfillmentCenters_path
    else
      render 'new'
    end
  end
 
  def update
    @fulfillmentCenter = FulfillmentCenter.find(params[:id])
 
    if @fulfillmentCenter.update(fulfillmentCenter_params)
      redirect_to fulfillmentCenters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @fulfillmentCenter = FulfillmentCenter.find(params[:id])
    @fulfillmentCenter.destroy
    redirect_to fulfillmentCenters_path
  end

 
  private
    def fulfillmentCenter_params
      params.require(:fulfillmentCenter).permit(:name, :centerCode, :address, :timezone, :asActive)
    end
end