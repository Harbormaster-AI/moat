class CardTokenizationsController < ApplicationController
  def index
    @cardTokenizations = CardTokenization.all
  end
 
  def show
    @cardTokenization = CardTokenization.find(params[:id])
  end
 
  def new
    @cardTokenization = CardTokenization.new
  end
 
  def edit
    @cardTokenization = CardTokenization.find(params[:id])
  end
 
  def create
    @cardTokenization = CardTokenization.new(cardTokenization_params)
 
    if @cardTokenization.save
      redirect_to cardTokenizations_path
    else
      render 'new'
    end
  end
 
  def update
    @cardTokenization = CardTokenization.find(params[:id])
 
    if @cardTokenization.update(cardTokenization_params)
      redirect_to cardTokenizations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @cardTokenization = CardTokenization.find(params[:id])
    @cardTokenization.destroy
    redirect_to cardTokenizations_path
  end

 
  private
    def cardTokenization_params
      params.require(:cardTokenization).permit(:tokenReference, :createdAt, :WalletProvider, :Status)
    end
end