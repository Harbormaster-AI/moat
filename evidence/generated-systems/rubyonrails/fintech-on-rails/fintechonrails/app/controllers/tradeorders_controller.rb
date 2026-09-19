class TradeOrdersController < ApplicationController
  def index
    @tradeOrders = TradeOrder.all
  end
 
  def show
    @tradeOrder = TradeOrder.find(params[:id])
  end
 
  def new
    @tradeOrder = TradeOrder.new
  end
 
  def edit
    @tradeOrder = TradeOrder.find(params[:id])
  end
 
  def create
    @tradeOrder = TradeOrder.new(tradeOrder_params)
 
    if @tradeOrder.save
      redirect_to tradeOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @tradeOrder = TradeOrder.find(params[:id])
 
    if @tradeOrder.update(tradeOrder_params)
      redirect_to tradeOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @tradeOrder = TradeOrder.find(params[:id])
    @tradeOrder.destroy
    redirect_to tradeOrders_path
  end

 
  private
    def tradeOrder_params
      params.require(:tradeOrder).permit(:orderId, :quantity, :limitPrice, :placedAt, :Side, :Type, :Status, :TimeInForce)
    end
end