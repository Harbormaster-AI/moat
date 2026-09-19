class ThirdPartysController < ApplicationController
  def index
    @thirdPartys = ThirdParty.all
  end
 
  def show
    @thirdParty = ThirdParty.find(params[:id])
  end
 
  def new
    @thirdParty = ThirdParty.new
  end
 
  def edit
    @thirdParty = ThirdParty.find(params[:id])
  end
 
  def create
    @thirdParty = ThirdParty.new(thirdParty_params)
 
    if @thirdParty.save
      redirect_to thirdPartys_path
    else
      render 'new'
    end
  end
 
  def update
    @thirdParty = ThirdParty.find(params[:id])
 
    if @thirdParty.update(thirdParty_params)
      redirect_to thirdPartys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @thirdParty = ThirdParty.find(params[:id])
    @thirdParty.destroy
    redirect_to thirdPartys_path
  end

 
  private
    def thirdParty_params
      params.require(:thirdParty).permit(:name, :taxId, :address, :PartyType)
    end
end