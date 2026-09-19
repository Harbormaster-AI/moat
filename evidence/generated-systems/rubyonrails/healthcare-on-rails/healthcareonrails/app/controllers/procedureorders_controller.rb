class ProcedureOrdersController < ApplicationController
  def index
    @procedureOrders = ProcedureOrder.all
  end
 
  def show
    @procedureOrder = ProcedureOrder.find(params[:id])
  end
 
  def new
    @procedureOrder = ProcedureOrder.new
  end
 
  def edit
    @procedureOrder = ProcedureOrder.find(params[:id])
  end
 
  def create
    @procedureOrder = ProcedureOrder.new(procedureOrder_params)
 
    if @procedureOrder.save
      redirect_to procedureOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @procedureOrder = ProcedureOrder.find(params[:id])
 
    if @procedureOrder.update(procedureOrder_params)
      redirect_to procedureOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @procedureOrder = ProcedureOrder.find(params[:id])
    @procedureOrder.destroy
    redirect_to procedureOrders_path
  end

 
  private
    def procedureOrder_params
      params.require(:procedureOrder).permit(:procedureCode, :consentObtained, :AnesthesiaType)
    end
end