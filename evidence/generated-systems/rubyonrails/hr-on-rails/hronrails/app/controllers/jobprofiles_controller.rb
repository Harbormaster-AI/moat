class JobProfilesController < ApplicationController
  def index
    @jobProfiles = JobProfile.all
  end
 
  def show
    @jobProfile = JobProfile.find(params[:id])
  end
 
  def new
    @jobProfile = JobProfile.new
  end
 
  def edit
    @jobProfile = JobProfile.find(params[:id])
  end
 
  def create
    @jobProfile = JobProfile.new(jobProfile_params)
 
    if @jobProfile.save
      redirect_to jobProfiles_path
    else
      render 'new'
    end
  end
 
  def update
    @jobProfile = JobProfile.find(params[:id])
 
    if @jobProfile.update(jobProfile_params)
      redirect_to jobProfiles_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @jobProfile = JobProfile.find(params[:id])
    @jobProfile.destroy
    redirect_to jobProfiles_path
  end

 
  private
    def jobProfile_params
      params.require(:jobProfile).permit(:title, :jobCode, :JobLevel, :ExemptStatus)
    end
end