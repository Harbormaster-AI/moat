class OrderItemsController < ApplicationController
  def index
    @orderItems = OrderItem.all
  end
 
  def show
    @orderItem = OrderItem.find(params[:id])
  end
 
  def new
    @orderItem = OrderItem.new
  end
 
  def edit
    @orderItem = OrderItem.find(params[:id])
  end
 
  def create
    @orderItem = OrderItem.new(orderItem_params)
 
    if @orderItem.save
      redirect_to orderItems_path
    else
      render 'new'
    end
  end
 
  def update
    @orderItem = OrderItem.find(params[:id])
 
    if @orderItem.update(orderItem_params)
      redirect_to orderItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @orderItem = OrderItem.find(params[:id])
    @orderItem.destroy
    redirect_to orderItems_path
  end

 
  private
    def orderItem_params
      params.require(:orderItem).permit(:quantity, :unitPrice, :discountAmount, :taxAmount, :totalAmount)
    end
end