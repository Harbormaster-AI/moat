class SalesOrderLinesController < ApplicationController
  def index
    @salesOrderLines = SalesOrderLine.all
  end
 
  def show
    @salesOrderLine = SalesOrderLine.find(params[:id])
  end
 
  def new
    @salesOrderLine = SalesOrderLine.new
  end
 
  def edit
    @salesOrderLine = SalesOrderLine.find(params[:id])
  end
 
  def create
    @salesOrderLine = SalesOrderLine.new(salesOrderLine_params)
 
    if @salesOrderLine.save
      redirect_to salesOrderLines_path
    else
      render 'new'
    end
  end
 
  def update
    @salesOrderLine = SalesOrderLine.find(params[:id])
 
    if @salesOrderLine.update(salesOrderLine_params)
      redirect_to salesOrderLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @salesOrderLine = SalesOrderLine.find(params[:id])
    @salesOrderLine.destroy
    redirect_to salesOrderLines_path
  end

 
  private
    def salesOrderLine_params
      params.require(:salesOrderLine).permit(:lineNumber, :quantity, :unitPrice, :dueDate)
    end
end