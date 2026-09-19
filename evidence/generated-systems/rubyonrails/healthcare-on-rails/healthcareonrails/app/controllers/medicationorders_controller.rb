class MedicationOrdersController < ApplicationController
  def index
    @medicationOrders = MedicationOrder.all
  end
 
  def show
    @medicationOrder = MedicationOrder.find(params[:id])
  end
 
  def new
    @medicationOrder = MedicationOrder.new
  end
 
  def edit
    @medicationOrder = MedicationOrder.find(params[:id])
  end
 
  def create
    @medicationOrder = MedicationOrder.new(medicationOrder_params)
 
    if @medicationOrder.save
      redirect_to medicationOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @medicationOrder = MedicationOrder.find(params[:id])
 
    if @medicationOrder.update(medicationOrder_params)
      redirect_to medicationOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @medicationOrder = MedicationOrder.find(params[:id])
    @medicationOrder.destroy
    redirect_to medicationOrders_path
  end

 
  private
    def medicationOrder_params
      params.require(:medicationOrder).permit(:medicationCode, :dose, :frequency, :duration, :Route)
    end
end