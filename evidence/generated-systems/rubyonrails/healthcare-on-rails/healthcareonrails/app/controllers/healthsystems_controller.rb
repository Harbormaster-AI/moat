class HealthSystemsController < ApplicationController
  def index
    @healthSystems = HealthSystem.all
  end
 
  def show
    @healthSystem = HealthSystem.find(params[:id])
  end
 
  def new
    @healthSystem = HealthSystem.new
  end
 
  def edit
    @healthSystem = HealthSystem.find(params[:id])
  end
 
  def create
    @healthSystem = HealthSystem.new(healthSystem_params)
 
    if @healthSystem.save
      redirect_to healthSystems_path
    else
      render 'new'
    end
  end
 
  def update
    @healthSystem = HealthSystem.find(params[:id])
 
    if @healthSystem.update(healthSystem_params)
      redirect_to healthSystems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @healthSystem = HealthSystem.find(params[:id])
    @healthSystem.destroy
    redirect_to healthSystems_path
  end

 
  private
    def healthSystem_params
      params.require(:healthSystem).permit(:name, :legalName, :headquartersCountry, :website)
    end
end