class ThirdPartyAssessmentsController < ApplicationController
  def index
    @thirdPartyAssessments = ThirdPartyAssessment.all
  end
 
  def show
    @thirdPartyAssessment = ThirdPartyAssessment.find(params[:id])
  end
 
  def new
    @thirdPartyAssessment = ThirdPartyAssessment.new
  end
 
  def edit
    @thirdPartyAssessment = ThirdPartyAssessment.find(params[:id])
  end
 
  def create
    @thirdPartyAssessment = ThirdPartyAssessment.new(thirdPartyAssessment_params)
 
    if @thirdPartyAssessment.save
      redirect_to thirdPartyAssessments_path
    else
      render 'new'
    end
  end
 
  def update
    @thirdPartyAssessment = ThirdPartyAssessment.find(params[:id])
 
    if @thirdPartyAssessment.update(thirdPartyAssessment_params)
      redirect_to thirdPartyAssessments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @thirdPartyAssessment = ThirdPartyAssessment.find(params[:id])
    @thirdPartyAssessment.destroy
    redirect_to thirdPartyAssessments_path
  end

 
  private
    def thirdPartyAssessment_params
      params.require(:thirdPartyAssessment).permit(:assessmentDate, :assessor, :AssessmentType, :Result)
    end
end