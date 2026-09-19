class OperatorsController < ApplicationController
  def index
    @operators = Operator.all
  end
 
  def show
    @operator = Operator.find(params[:id])
  end
 
  def new
    @operator = Operator.new
  end
 
  def edit
    @operator = Operator.find(params[:id])
  end
 
  def create
    @operator = Operator.new(operator_params)
 
    if @operator.save
      redirect_to operators_path
    else
      render 'new'
    end
  end
 
  def update
    @operator = Operator.find(params[:id])
 
    if @operator.update(operator_params)
      redirect_to operators_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @operator = Operator.find(params[:id])
    @operator.destroy
    redirect_to operators_path
  end

 
  private
    def operator_params
      params.require(:operator).permit(:name, :icaoDesignator, :OperatorType)
    end
end