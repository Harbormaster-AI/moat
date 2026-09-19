class SoftwareUpdatesController < ApplicationController
  def index
    @softwareUpdates = SoftwareUpdate.all
  end
 
  def show
    @softwareUpdate = SoftwareUpdate.find(params[:id])
  end
 
  def new
    @softwareUpdate = SoftwareUpdate.new
  end
 
  def edit
    @softwareUpdate = SoftwareUpdate.find(params[:id])
  end
 
  def create
    @softwareUpdate = SoftwareUpdate.new(softwareUpdate_params)
 
    if @softwareUpdate.save
      redirect_to softwareUpdates_path
    else
      render 'new'
    end
  end
 
  def update
    @softwareUpdate = SoftwareUpdate.find(params[:id])
 
    if @softwareUpdate.update(softwareUpdate_params)
      redirect_to softwareUpdates_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @softwareUpdate = SoftwareUpdate.find(params[:id])
    @softwareUpdate.destroy
    redirect_to softwareUpdates_path
  end

 
  private
    def softwareUpdate_params
      params.require(:softwareUpdate).permit(:version, :appliedDate, :UpdateType)
    end
end