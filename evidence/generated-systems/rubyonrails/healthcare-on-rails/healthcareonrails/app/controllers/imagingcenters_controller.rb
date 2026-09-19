class ImagingCentersController < ApplicationController
  def index
    @imagingCenters = ImagingCenter.all
  end
 
  def show
    @imagingCenter = ImagingCenter.find(params[:id])
  end
 
  def new
    @imagingCenter = ImagingCenter.new
  end
 
  def edit
    @imagingCenter = ImagingCenter.find(params[:id])
  end
 
  def create
    @imagingCenter = ImagingCenter.new(imagingCenter_params)
 
    if @imagingCenter.save
      redirect_to imagingCenters_path
    else
      render 'new'
    end
  end
 
  def update
    @imagingCenter = ImagingCenter.find(params[:id])
 
    if @imagingCenter.update(imagingCenter_params)
      redirect_to imagingCenters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @imagingCenter = ImagingCenter.find(params[:id])
    @imagingCenter.destroy
    redirect_to imagingCenters_path
  end

 
  private
    def imagingCenter_params
      params.require(:imagingCenter).permit(:name)
    end
end