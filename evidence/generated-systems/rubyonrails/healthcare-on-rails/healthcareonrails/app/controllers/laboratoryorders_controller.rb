class LaboratoryOrdersController < ApplicationController
  def index
    @laboratoryOrders = LaboratoryOrder.all
  end
 
  def show
    @laboratoryOrder = LaboratoryOrder.find(params[:id])
  end
 
  def new
    @laboratoryOrder = LaboratoryOrder.new
  end
 
  def edit
    @laboratoryOrder = LaboratoryOrder.find(params[:id])
  end
 
  def create
    @laboratoryOrder = LaboratoryOrder.new(laboratoryOrder_params)
 
    if @laboratoryOrder.save
      redirect_to laboratoryOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @laboratoryOrder = LaboratoryOrder.find(params[:id])
 
    if @laboratoryOrder.update(laboratoryOrder_params)
      redirect_to laboratoryOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @laboratoryOrder = LaboratoryOrder.find(params[:id])
    @laboratoryOrder.destroy
    redirect_to laboratoryOrders_path
  end

 
  private
    def laboratoryOrder_params
      params.require(:laboratoryOrder).permit(:testCode, :fastingRequired, :SpecimenType)
    end
end