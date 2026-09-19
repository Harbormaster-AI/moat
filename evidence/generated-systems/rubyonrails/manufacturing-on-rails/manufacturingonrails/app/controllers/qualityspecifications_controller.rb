class QualitySpecificationsController < ApplicationController
  def index
    @qualitySpecifications = QualitySpecification.all
  end
 
  def show
    @qualitySpecification = QualitySpecification.find(params[:id])
  end
 
  def new
    @qualitySpecification = QualitySpecification.new
  end
 
  def edit
    @qualitySpecification = QualitySpecification.find(params[:id])
  end
 
  def create
    @qualitySpecification = QualitySpecification.new(qualitySpecification_params)
 
    if @qualitySpecification.save
      redirect_to qualitySpecifications_path
    else
      render 'new'
    end
  end
 
  def update
    @qualitySpecification = QualitySpecification.find(params[:id])
 
    if @qualitySpecification.update(qualitySpecification_params)
      redirect_to qualitySpecifications_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @qualitySpecification = QualitySpecification.find(params[:id])
    @qualitySpecification.destroy
    redirect_to qualitySpecifications_path
  end

 
  private
    def qualitySpecification_params
      params.require(:qualitySpecification).permit(:specCode, :name, :version)
    end
end