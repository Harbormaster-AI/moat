class OrderLinesController < ApplicationController
  def index
    @orderLines = OrderLine.all
  end
 
  def show
    @orderLine = OrderLine.find(params[:id])
  end
 
  def new
    @orderLine = OrderLine.new
  end
 
  def edit
    @orderLine = OrderLine.find(params[:id])
  end
 
  def create
    @orderLine = OrderLine.new(orderLine_params)
 
    if @orderLine.save
      redirect_to orderLines_path
    else
      render 'new'
    end
  end
 
  def update
    @orderLine = OrderLine.find(params[:id])
 
    if @orderLine.update(orderLine_params)
      redirect_to orderLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @orderLine = OrderLine.find(params[:id])
    @orderLine.destroy
    redirect_to orderLines_path
  end

 
  private
    def orderLine_params
      params.require(:orderLine).permit(:quantity, :unitPrice, :totalPrice, :taxRate, :LineStatus)
    end
end