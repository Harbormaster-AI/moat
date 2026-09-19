class LabResultsController < ApplicationController
  def index
    @labResults = LabResult.all
  end
 
  def show
    @labResult = LabResult.find(params[:id])
  end
 
  def new
    @labResult = LabResult.new
  end
 
  def edit
    @labResult = LabResult.find(params[:id])
  end
 
  def create
    @labResult = LabResult.new(labResult_params)
 
    if @labResult.save
      redirect_to labResults_path
    else
      render 'new'
    end
  end
 
  def update
    @labResult = LabResult.find(params[:id])
 
    if @labResult.update(labResult_params)
      redirect_to labResults_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @labResult = LabResult.find(params[:id])
    @labResult.destroy
    redirect_to labResults_path
  end

 
  private
    def labResult_params
      params.require(:labResult).permit(:resultCode, :issuedDate, :Status)
    end
end