class ReturnRequestsController < ApplicationController
  def index
    @returnRequests = ReturnRequest.all
  end
 
  def show
    @returnRequest = ReturnRequest.find(params[:id])
  end
 
  def new
    @returnRequest = ReturnRequest.new
  end
 
  def edit
    @returnRequest = ReturnRequest.find(params[:id])
  end
 
  def create
    @returnRequest = ReturnRequest.new(returnRequest_params)
 
    if @returnRequest.save
      redirect_to returnRequests_path
    else
      render 'new'
    end
  end
 
  def update
    @returnRequest = ReturnRequest.find(params[:id])
 
    if @returnRequest.update(returnRequest_params)
      redirect_to returnRequests_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @returnRequest = ReturnRequest.find(params[:id])
    @returnRequest.destroy
    redirect_to returnRequests_path
  end

 
  private
    def returnRequest_params
      params.require(:returnRequest).permit(:returnNumber, :createdAt, :refundAmount, :Status)
    end
end