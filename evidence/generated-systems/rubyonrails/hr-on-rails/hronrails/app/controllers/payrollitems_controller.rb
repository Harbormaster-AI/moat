class PayrollItemsController < ApplicationController
  def index
    @payrollItems = PayrollItem.all
  end
 
  def show
    @payrollItem = PayrollItem.find(params[:id])
  end
 
  def new
    @payrollItem = PayrollItem.new
  end
 
  def edit
    @payrollItem = PayrollItem.find(params[:id])
  end
 
  def create
    @payrollItem = PayrollItem.new(payrollItem_params)
 
    if @payrollItem.save
      redirect_to payrollItems_path
    else
      render 'new'
    end
  end
 
  def update
    @payrollItem = PayrollItem.find(params[:id])
 
    if @payrollItem.update(payrollItem_params)
      redirect_to payrollItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @payrollItem = PayrollItem.find(params[:id])
    @payrollItem.destroy
    redirect_to payrollItems_path
  end

 
  private
    def payrollItem_params
      params.require(:payrollItem).permit(:amount, :taxable, :ItemType)
    end
end