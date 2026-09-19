class BenefitEnrollmentsController < ApplicationController
  def index
    @benefitEnrollments = BenefitEnrollment.all
  end
 
  def show
    @benefitEnrollment = BenefitEnrollment.find(params[:id])
  end
 
  def new
    @benefitEnrollment = BenefitEnrollment.new
  end
 
  def edit
    @benefitEnrollment = BenefitEnrollment.find(params[:id])
  end
 
  def create
    @benefitEnrollment = BenefitEnrollment.new(benefitEnrollment_params)
 
    if @benefitEnrollment.save
      redirect_to benefitEnrollments_path
    else
      render 'new'
    end
  end
 
  def update
    @benefitEnrollment = BenefitEnrollment.find(params[:id])
 
    if @benefitEnrollment.update(benefitEnrollment_params)
      redirect_to benefitEnrollments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @benefitEnrollment = BenefitEnrollment.find(params[:id])
    @benefitEnrollment.destroy
    redirect_to benefitEnrollments_path
  end

 
  private
    def benefitEnrollment_params
      params.require(:benefitEnrollment).permit(:enrollmentId, :effectiveFrom, :effectiveTo, :Status, :CoverageLevel)
    end
end