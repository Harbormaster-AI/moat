class SecuritysController < ApplicationController
  def index
    @securitys = Security.all
  end
 
  def show
    @security = Security.find(params[:id])
  end
 
  def new
    @security = Security.new
  end
 
  def edit
    @security = Security.find(params[:id])
  end
 
  def create
    @security = Security.new(security_params)
 
    if @security.save
      redirect_to securitys_path
    else
      render 'new'
    end
  end
 
  def update
    @security = Security.find(params[:id])
 
    if @security.update(security_params)
      redirect_to securitys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @security = Security.find(params[:id])
    @security.destroy
    redirect_to securitys_path
  end

 
  private
    def security_params
      params.require(:security).permit(:symbol, :isin, :cusip, :currency, :SecurityType)
    end
end