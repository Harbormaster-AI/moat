class ReturnItemsController < ApplicationController
  def index
    @returnItems = ReturnItem.all
  end
 
  def show
    @returnItem = ReturnItem.find(params[:id])
  end
 
  def new
    @returnItem = ReturnItem.new
  end
 
  def edit
    @returnItem = ReturnItem.find(params[:id])
  end
 
  def create
    @returnItem = ReturnItem.new(returnItem_params)
 
    if @returnItem.save
      redirect_to returnItems_path
    else
      render 'new'
    end
  end
 
  def update
    @returnItem = ReturnItem.find(params[:id])
 
    if @returnItem.update(returnItem_params)
      redirect_to returnItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @returnItem = ReturnItem.find(params[:id])
    @returnItem.destroy
    redirect_to returnItems_path
  end

 
  private
    def returnItem_params
      params.require(:returnItem).permit(:quantity, :Reason, :Condition)
    end
end