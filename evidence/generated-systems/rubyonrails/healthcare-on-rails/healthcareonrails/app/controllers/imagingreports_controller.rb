class ImagingReportsController < ApplicationController
  def index
    @imagingReports = ImagingReport.all
  end
 
  def show
    @imagingReport = ImagingReport.find(params[:id])
  end
 
  def new
    @imagingReport = ImagingReport.new
  end
 
  def edit
    @imagingReport = ImagingReport.find(params[:id])
  end
 
  def create
    @imagingReport = ImagingReport.new(imagingReport_params)
 
    if @imagingReport.save
      redirect_to imagingReports_path
    else
      render 'new'
    end
  end
 
  def update
    @imagingReport = ImagingReport.find(params[:id])
 
    if @imagingReport.update(imagingReport_params)
      redirect_to imagingReports_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @imagingReport = ImagingReport.find(params[:id])
    @imagingReport.destroy
    redirect_to imagingReports_path
  end

 
  private
    def imagingReport_params
      params.require(:imagingReport).permit(:reportNumber, :impression, :reportedDate, :Status)
    end
end