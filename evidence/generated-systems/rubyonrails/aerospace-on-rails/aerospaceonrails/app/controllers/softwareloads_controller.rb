class SoftwareLoadsController < ApplicationController
  def index
    @softwareLoads = SoftwareLoad.all
  end
 
  def show
    @softwareLoad = SoftwareLoad.find(params[:id])
  end
 
  def new
    @softwareLoad = SoftwareLoad.new
  end
 
  def edit
    @softwareLoad = SoftwareLoad.find(params[:id])
  end
 
  def create
    @softwareLoad = SoftwareLoad.new(softwareLoad_params)
 
    if @softwareLoad.save
      redirect_to softwareLoads_path
    else
      render 'new'
    end
  end
 
  def update
    @softwareLoad = SoftwareLoad.find(params[:id])
 
    if @softwareLoad.update(softwareLoad_params)
      redirect_to softwareLoads_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @softwareLoad = SoftwareLoad.find(params[:id])
    @softwareLoad.destroy
    redirect_to softwareLoads_path
  end

 
  private
    def softwareLoad_params
      params.require(:softwareLoad).permit(:version, :LoadType)
    end
end