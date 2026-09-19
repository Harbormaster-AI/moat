class CustomerAddresssController < ApplicationController
  def index
    @customerAddresss = CustomerAddress.all
  end
 
  def show
    @customerAddress = CustomerAddress.find(params[:id])
  end
 
  def new
    @customerAddress = CustomerAddress.new
  end
 
  def edit
    @customerAddress = CustomerAddress.find(params[:id])
  end
 
  def create
    @customerAddress = CustomerAddress.new(customerAddress_params)
 
    if @customerAddress.save
      redirect_to customerAddresss_path
    else
      render 'new'
    end
  end
 
  def update
    @customerAddress = CustomerAddress.find(params[:id])
 
    if @customerAddress.update(customerAddress_params)
      redirect_to customerAddresss_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @customerAddress = CustomerAddress.find(params[:id])
    @customerAddress.destroy
    redirect_to customerAddresss_path
  end

 
  private
    def customerAddress_params
      params.require(:customerAddress).permit(:label, :address, :asDefaultShipping, :asDefaultBilling)
    end
end