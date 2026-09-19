class MRPRunsController < ApplicationController
  def index
    @mRPRuns = MRPRun.all
  end
 
  def show
    @mRPRun = MRPRun.find(params[:id])
  end
 
  def new
    @mRPRun = MRPRun.new
  end
 
  def edit
    @mRPRun = MRPRun.find(params[:id])
  end
 
  def create
    @mRPRun = MRPRun.new(mRPRun_params)
 
    if @mRPRun.save
      redirect_to mRPRuns_path
    else
      render 'new'
    end
  end
 
  def update
    @mRPRun = MRPRun.find(params[:id])
 
    if @mRPRun.update(mRPRun_params)
      redirect_to mRPRuns_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @mRPRun = MRPRun.find(params[:id])
    @mRPRun.destroy
    redirect_to mRPRuns_path
  end

 
  private
    def mRPRun_params
      params.require(:mRPRun).permit(:runNumber, :runDateTime, :planningHorizonDays, :Status)
    end
end