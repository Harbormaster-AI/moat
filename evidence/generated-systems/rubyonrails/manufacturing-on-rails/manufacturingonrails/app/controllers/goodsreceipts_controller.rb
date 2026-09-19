class GoodsReceiptsController < ApplicationController
  def index
    @goodsReceipts = GoodsReceipt.all
  end
 
  def show
    @goodsReceipt = GoodsReceipt.find(params[:id])
  end
 
  def new
    @goodsReceipt = GoodsReceipt.new
  end
 
  def edit
    @goodsReceipt = GoodsReceipt.find(params[:id])
  end
 
  def create
    @goodsReceipt = GoodsReceipt.new(goodsReceipt_params)
 
    if @goodsReceipt.save
      redirect_to goodsReceipts_path
    else
      render 'new'
    end
  end
 
  def update
    @goodsReceipt = GoodsReceipt.find(params[:id])
 
    if @goodsReceipt.update(goodsReceipt_params)
      redirect_to goodsReceipts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @goodsReceipt = GoodsReceipt.find(params[:id])
    @goodsReceipt.destroy
    redirect_to goodsReceipts_path
  end

 
  private
    def goodsReceipt_params
      params.require(:goodsReceipt).permit(:receiptNumber, :receiptDate, :Status)
    end
end