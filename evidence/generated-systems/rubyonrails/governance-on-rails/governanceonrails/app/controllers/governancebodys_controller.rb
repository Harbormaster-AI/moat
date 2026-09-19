class GovernanceBodysController < ApplicationController
  def index
    @governanceBodys = GovernanceBody.all
  end
 
  def show
    @governanceBody = GovernanceBody.find(params[:id])
  end
 
  def new
    @governanceBody = GovernanceBody.new
  end
 
  def edit
    @governanceBody = GovernanceBody.find(params[:id])
  end
 
  def create
    @governanceBody = GovernanceBody.new(governanceBody_params)
 
    if @governanceBody.save
      redirect_to governanceBodys_path
    else
      render 'new'
    end
  end
 
  def update
    @governanceBody = GovernanceBody.find(params[:id])
 
    if @governanceBody.update(governanceBody_params)
      redirect_to governanceBodys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @governanceBody = GovernanceBody.find(params[:id])
    @governanceBody.destroy
    redirect_to governanceBodys_path
  end

 
  private
    def governanceBody_params
      params.require(:governanceBody).permit(:name, :charterUrl, :chair, :BodyType)
    end
end