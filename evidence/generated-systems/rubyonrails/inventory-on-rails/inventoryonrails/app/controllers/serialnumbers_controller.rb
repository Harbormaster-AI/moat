class SerialNumbersController < ApplicationController
  def index
    @serialNumbers = SerialNumber.all
  end
 
  def show
    @serialNumber = SerialNumber.find(params[:id])
  end
 
  def new
    @serialNumber = SerialNumber.new
  end
 
  def edit
    @serialNumber = SerialNumber.find(params[:id])
  end
 
  def create
    @serialNumber = SerialNumber.new(serialNumber_params)
 
    if @serialNumber.save
      redirect_to serialNumbers_path
    else
      render 'new'
    end
  end
 
  def update
    @serialNumber = SerialNumber.find(params[:id])
 
    if @serialNumber.update(serialNumber_params)
      redirect_to serialNumbers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @serialNumber = SerialNumber.find(params[:id])
    @serialNumber.destroy
    redirect_to serialNumbers_path
  end

 
  private
    def serialNumber_params
      params.require(:serialNumber).permit(:serial, :activationDate, :Status)
    end
end