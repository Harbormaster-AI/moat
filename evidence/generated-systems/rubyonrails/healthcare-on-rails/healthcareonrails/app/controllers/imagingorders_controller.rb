class ImagingOrdersController < ApplicationController
  def index
    @imagingOrders = ImagingOrder.all
  end
 
  def show
    @imagingOrder = ImagingOrder.find(params[:id])
  end
 
  def new
    @imagingOrder = ImagingOrder.new
  end
 
  def edit
    @imagingOrder = ImagingOrder.find(params[:id])
  end
 
  def create
    @imagingOrder = ImagingOrder.new(imagingOrder_params)
 
    if @imagingOrder.save
      redirect_to imagingOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @imagingOrder = ImagingOrder.find(params[:id])
 
    if @imagingOrder.update(imagingOrder_params)
      redirect_to imagingOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @imagingOrder = ImagingOrder.find(params[:id])
    @imagingOrder.destroy
    redirect_to imagingOrders_path
  end

 
  private
    def imagingOrder_params
      params.require(:imagingOrder).permit(:bodySite, :contrast, :Modality)
    end
end