class OperationsController < ApplicationController
  def index
    @operations = Operation.all
  end
 
  def show
    @operation = Operation.find(params[:id])
  end
 
  def new
    @operation = Operation.new
  end
 
  def edit
    @operation = Operation.find(params[:id])
  end
 
  def create
    @operation = Operation.new(operation_params)
 
    if @operation.save
      redirect_to operations_path
    else
      render 'new'
    end
  end
 
  def update
    @operation = Operation.find(params[:id])
 
    if @operation.update(operation_params)
      redirect_to operations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @operation = Operation.find(params[:id])
    @operation.destroy
    redirect_to operations_path
  end

 
  private
    def operation_params
      params.require(:operation).permit(:operationNumber, :name, :setupTime, :standardCycleTime, :OperationType)
    end
end