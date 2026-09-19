class ClinicalOrdersController < ApplicationController
  def index
    @clinicalOrders = ClinicalOrder.all
  end
 
  def show
    @clinicalOrder = ClinicalOrder.find(params[:id])
  end
 
  def new
    @clinicalOrder = ClinicalOrder.new
  end
 
  def edit
    @clinicalOrder = ClinicalOrder.find(params[:id])
  end
 
  def create
    @clinicalOrder = ClinicalOrder.new(clinicalOrder_params)
 
    if @clinicalOrder.save
      redirect_to clinicalOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @clinicalOrder = ClinicalOrder.find(params[:id])
 
    if @clinicalOrder.update(clinicalOrder_params)
      redirect_to clinicalOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @clinicalOrder = ClinicalOrder.find(params[:id])
    @clinicalOrder.destroy
    redirect_to clinicalOrders_path
  end

 
  private
    def clinicalOrder_params
      params.require(:clinicalOrder).permit(:orderNumber, :Status, :OrderType, :Priority)
    end
end