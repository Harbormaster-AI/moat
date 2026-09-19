class GoodsReceiptLinesController < ApplicationController
  def index
    @goodsReceiptLines = GoodsReceiptLine.all
  end
 
  def show
    @goodsReceiptLine = GoodsReceiptLine.find(params[:id])
  end
 
  def new
    @goodsReceiptLine = GoodsReceiptLine.new
  end
 
  def edit
    @goodsReceiptLine = GoodsReceiptLine.find(params[:id])
  end
 
  def create
    @goodsReceiptLine = GoodsReceiptLine.new(goodsReceiptLine_params)
 
    if @goodsReceiptLine.save
      redirect_to goodsReceiptLines_path
    else
      render 'new'
    end
  end
 
  def update
    @goodsReceiptLine = GoodsReceiptLine.find(params[:id])
 
    if @goodsReceiptLine.update(goodsReceiptLine_params)
      redirect_to goodsReceiptLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @goodsReceiptLine = GoodsReceiptLine.find(params[:id])
    @goodsReceiptLine.destroy
    redirect_to goodsReceiptLines_path
  end

 
  private
    def goodsReceiptLine_params
      params.require(:goodsReceiptLine).permit(:lineNumber, :receivedQuantity, :acceptedQuantity, :rejectedQuantity, :lot)
    end
end