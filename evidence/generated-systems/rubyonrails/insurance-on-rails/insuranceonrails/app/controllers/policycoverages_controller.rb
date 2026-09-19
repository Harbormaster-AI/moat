class PolicyCoveragesController < ApplicationController
  def index
    @policyCoverages = PolicyCoverage.all
  end
 
  def show
    @policyCoverage = PolicyCoverage.find(params[:id])
  end
 
  def new
    @policyCoverage = PolicyCoverage.new
  end
 
  def edit
    @policyCoverage = PolicyCoverage.find(params[:id])
  end
 
  def create
    @policyCoverage = PolicyCoverage.new(policyCoverage_params)
 
    if @policyCoverage.save
      redirect_to policyCoverages_path
    else
      render 'new'
    end
  end
 
  def update
    @policyCoverage = PolicyCoverage.find(params[:id])
 
    if @policyCoverage.update(policyCoverage_params)
      redirect_to policyCoverages_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @policyCoverage = PolicyCoverage.find(params[:id])
    @policyCoverage.destroy
    redirect_to policyCoverages_path
  end

 
  private
    def policyCoverage_params
      params.require(:policyCoverage).permit(:limit, :deductible, :premium, :CoverageType)
    end
end