class CoveragesController < ApplicationController
  def index
    @coverages = Coverage.all
  end
 
  def show
    @coverage = Coverage.find(params[:id])
  end
 
  def new
    @coverage = Coverage.new
  end
 
  def edit
    @coverage = Coverage.find(params[:id])
  end
 
  def create
    @coverage = Coverage.new(coverage_params)
 
    if @coverage.save
      redirect_to coverages_path
    else
      render 'new'
    end
  end
 
  def update
    @coverage = Coverage.find(params[:id])
 
    if @coverage.update(coverage_params)
      redirect_to coverages_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @coverage = Coverage.find(params[:id])
    @coverage.destroy
    redirect_to coverages_path
  end

 
  private
    def coverage_params
      params.require(:coverage).permit(:memberId, :groupNumber, :effectiveDate, :endDate, :CoverageType)
    end
end