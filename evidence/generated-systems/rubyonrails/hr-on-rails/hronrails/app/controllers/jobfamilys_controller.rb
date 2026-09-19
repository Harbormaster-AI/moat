class JobFamilysController < ApplicationController
  def index
    @jobFamilys = JobFamily.all
  end
 
  def show
    @jobFamily = JobFamily.find(params[:id])
  end
 
  def new
    @jobFamily = JobFamily.new
  end
 
  def edit
    @jobFamily = JobFamily.find(params[:id])
  end
 
  def create
    @jobFamily = JobFamily.new(jobFamily_params)
 
    if @jobFamily.save
      redirect_to jobFamilys_path
    else
      render 'new'
    end
  end
 
  def update
    @jobFamily = JobFamily.find(params[:id])
 
    if @jobFamily.update(jobFamily_params)
      redirect_to jobFamilys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @jobFamily = JobFamily.find(params[:id])
    @jobFamily.destroy
    redirect_to jobFamilys_path
  end

 
  private
    def jobFamily_params
      params.require(:jobFamily).permit(:name, :description)
    end
end