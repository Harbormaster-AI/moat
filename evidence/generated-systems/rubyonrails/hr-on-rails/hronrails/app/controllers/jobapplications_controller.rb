class JobApplicationsController < ApplicationController
  def index
    @jobApplications = JobApplication.all
  end
 
  def show
    @jobApplication = JobApplication.find(params[:id])
  end
 
  def new
    @jobApplication = JobApplication.new
  end
 
  def edit
    @jobApplication = JobApplication.find(params[:id])
  end
 
  def create
    @jobApplication = JobApplication.new(jobApplication_params)
 
    if @jobApplication.save
      redirect_to jobApplications_path
    else
      render 'new'
    end
  end
 
  def update
    @jobApplication = JobApplication.find(params[:id])
 
    if @jobApplication.update(jobApplication_params)
      redirect_to jobApplications_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @jobApplication = JobApplication.find(params[:id])
    @jobApplication.destroy
    redirect_to jobApplications_path
  end

 
  private
    def jobApplication_params
      params.require(:jobApplication).permit(:applicationNumber, :appliedDate, :resumeUrl, :Status)
    end
end