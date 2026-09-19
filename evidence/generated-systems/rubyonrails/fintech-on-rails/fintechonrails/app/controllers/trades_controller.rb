class TradesController < ApplicationController
  def index
    @trades = Trade.all
  end
 
  def show
    @trade = Trade.find(params[:id])
  end
 
  def new
    @trade = Trade.new
  end
 
  def edit
    @trade = Trade.find(params[:id])
  end
 
  def create
    @trade = Trade.new(trade_params)
 
    if @trade.save
      redirect_to trades_path
    else
      render 'new'
    end
  end
 
  def update
    @trade = Trade.find(params[:id])
 
    if @trade.update(trade_params)
      redirect_to trades_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @trade = Trade.find(params[:id])
    @trade.destroy
    redirect_to trades_path
  end

 
  private
    def trade_params
      params.require(:trade).permit(:executedAt, :quantity, :price, :fees, :settlementDate)
    end
end