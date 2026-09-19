class APIClientsController < ApplicationController
  def index
    @aPIClients = APIClient.all
  end
 
  def show
    @aPIClient = APIClient.find(params[:id])
  end
 
  def new
    @aPIClient = APIClient.new
  end
 
  def edit
    @aPIClient = APIClient.find(params[:id])
  end
 
  def create
    @aPIClient = APIClient.new(aPIClient_params)
 
    if @aPIClient.save
      redirect_to aPIClients_path
    else
      render 'new'
    end
  end
 
  def update
    @aPIClient = APIClient.find(params[:id])
 
    if @aPIClient.update(aPIClient_params)
      redirect_to aPIClients_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aPIClient = APIClient.find(params[:id])
    @aPIClient.destroy
    redirect_to aPIClients_path
  end

 
  private
    def aPIClient_params
      params.require(:aPIClient).permit(:name, :clientId, :redirectUri, :ClientType)
    end
end