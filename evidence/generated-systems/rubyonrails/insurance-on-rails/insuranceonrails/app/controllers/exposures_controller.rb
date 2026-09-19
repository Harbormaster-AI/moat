class ExposuresController < ApplicationController
  def index
    @exposures = Exposure.all
  end
 
  def show
    @exposure = Exposure.find(params[:id])
  end
 
  def new
    @exposure = Exposure.new
  end
 
  def edit
    @exposure = Exposure.find(params[:id])
  end
 
  def create
    @exposure = Exposure.new(exposure_params)
 
    if @exposure.save
      redirect_to exposures_path
    else
      render 'new'
    end
  end
 
  def update
    @exposure = Exposure.find(params[:id])
 
    if @exposure.update(exposure_params)
      redirect_to exposures_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @exposure = Exposure.find(params[:id])
    @exposure.destroy
    redirect_to exposures_path
  end

 
  private
    def exposure_params
      params.require(:exposure).permit(:ExposureType, :Status)
    end
end