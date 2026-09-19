class JobRequisitionsController < ApplicationController
  def index
    @jobRequisitions = JobRequisition.all
  end
 
  def show
    @jobRequisition = JobRequisition.find(params[:id])
  end
 
  def new
    @jobRequisition = JobRequisition.new
  end
 
  def edit
    @jobRequisition = JobRequisition.find(params[:id])
  end
 
  def create
    @jobRequisition = JobRequisition.new(jobRequisition_params)
 
    if @jobRequisition.save
      redirect_to jobRequisitions_path
    else
      render 'new'
    end
  end
 
  def update
    @jobRequisition = JobRequisition.find(params[:id])
 
    if @jobRequisition.update(jobRequisition_params)
      redirect_to jobRequisitions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @jobRequisition = JobRequisition.find(params[:id])
    @jobRequisition.destroy
    redirect_to jobRequisitions_path
  end

 
  private
    def jobRequisition_params
      params.require(:jobRequisition).permit(:requisitionNumber, :title, :openings, :targetStartDate, :Status, :Priority)
    end
end