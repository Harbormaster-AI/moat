class ServiceProvidersController < ApplicationController
  def index
    @serviceProviders = ServiceProvider.all
  end
 
  def show
    @serviceProvider = ServiceProvider.find(params[:id])
  end
 
  def new
    @serviceProvider = ServiceProvider.new
  end
 
  def edit
    @serviceProvider = ServiceProvider.find(params[:id])
  end
 
  def create
    @serviceProvider = ServiceProvider.new(serviceProvider_params)
 
    if @serviceProvider.save
      redirect_to serviceProviders_path
    else
      render 'new'
    end
  end
 
  def update
    @serviceProvider = ServiceProvider.find(params[:id])
 
    if @serviceProvider.update(serviceProvider_params)
      redirect_to serviceProviders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @serviceProvider = ServiceProvider.find(params[:id])
    @serviceProvider.destroy
    redirect_to serviceProviders_path
  end

 
  private
    def serviceProvider_params
      params.require(:serviceProvider).permit(:name, :taxId, :ProviderType, :NetworkStatus)
    end
end