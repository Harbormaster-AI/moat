class OutboundAllocationsController < ApplicationController
  def index
    @outboundAllocations = OutboundAllocation.all
  end
 
  def show
    @outboundAllocation = OutboundAllocation.find(params[:id])
  end
 
  def new
    @outboundAllocation = OutboundAllocation.new
  end
 
  def edit
    @outboundAllocation = OutboundAllocation.find(params[:id])
  end
 
  def create
    @outboundAllocation = OutboundAllocation.new(outboundAllocation_params)
 
    if @outboundAllocation.save
      redirect_to outboundAllocations_path
    else
      render 'new'
    end
  end
 
  def update
    @outboundAllocation = OutboundAllocation.find(params[:id])
 
    if @outboundAllocation.update(outboundAllocation_params)
      redirect_to outboundAllocations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @outboundAllocation = OutboundAllocation.find(params[:id])
    @outboundAllocation.destroy
    redirect_to outboundAllocations_path
  end

 
  private
    def outboundAllocation_params
      params.require(:outboundAllocation).permit(:allocationNumber, :allocatedQuantity, :allocationDate, :Status)
    end
end