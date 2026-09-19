class QualityChecksController < ApplicationController
  def index
    @qualityChecks = QualityCheck.all
  end
 
  def show
    @qualityCheck = QualityCheck.find(params[:id])
  end
 
  def new
    @qualityCheck = QualityCheck.new
  end
 
  def edit
    @qualityCheck = QualityCheck.find(params[:id])
  end
 
  def create
    @qualityCheck = QualityCheck.new(qualityCheck_params)
 
    if @qualityCheck.save
      redirect_to qualityChecks_path
    else
      render 'new'
    end
  end
 
  def update
    @qualityCheck = QualityCheck.find(params[:id])
 
    if @qualityCheck.update(qualityCheck_params)
      redirect_to qualityChecks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @qualityCheck = QualityCheck.find(params[:id])
    @qualityCheck.destroy
    redirect_to qualityChecks_path
  end

 
  private
    def qualityCheck_params
      params.require(:qualityCheck).permit(:checkedAt, :observedValue, :sampleSize, :Status)
    end
end